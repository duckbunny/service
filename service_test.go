package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

var TestStruct *Service
var TestJSON []byte

func init() {
	var err error
	TestJSON, err = ioutil.ReadFile("service.json")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, TestJSON); err != nil {
		log.Fatal(err)
	}
	TestJSON = buf.Bytes()
	TestStruct = &Service{
		Title:    "service",
		Domain:   "duckbunny",
		Version:  "0.1",
		Type:     "test",
		Protocol: "http",
		Private:  false,
		Method:   "POST",
		Requires: []Service{
			Service{
				Title:   "service2",
				Domain:  "duckbunny",
				Version: "0.1",
			},
			Service{
				Title:   "service3",
				Domain:  "duckbunny",
				Version: "0.1",
			},
		},
		Parameters: []Parameter{
			Parameter{
				Key:         "testparam1",
				Description: "My first test parameter",
				Required:    false,
				Type:        "json",
				DataType:    "map[string]string",
			},
			Parameter{
				Key:         "testparam2",
				Description: "My second test parameter",
				Required:    true,
				Type:        "header",
				DataType:    "string",
			},
		},
		Configs: []Config{
			Config{
				Key:         "testconfig1",
				Description: "My first test config",
				Required:    false,
			},
			Config{
				Key:         "testconfig2",
				Description: "My second test config",
				Required:    true,
			},
		},
		Response: Response{
			Type:     "googlejson",
			DataType: "map[string]string",
		},
		Flags: []Flag{
			Flag{
				Key:         "test",
				Env:         "TESTVAR",
				Description: "This is a test flag.",
				Required:    true,
			},
		},
		Port: "80",
		Host: "localhost",
	}
}

func TestThis(t *testing.T) {
	serviceFile = "fakefile.yaml"
	_, err := This()
	if err == nil {
		t.Error("Expected file load error")
	}
	serviceFile = "Service.yaml"
	_, err = This()
	if err != ErrNoPort {
		t.Error("Expected no port error")
	}
	servicePort = "80"
	_, err = This()
	if err != ErrNoHost {
		t.Error("Expected no host error")
	}
	serviceHost = "localhost"
	var s *Service
	s, err = This()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(s, TestStruct) {
		t.Error("Test Failed.")
	}
}

func TestLoadFromFile(t *testing.T) {
	_, err := LoadFromFile("fakefile.yaml")
	if err == nil {
		t.Error("Expected file load error")
	}
	TestStruct.Host = ""
	TestStruct.Port = ""
	s, err := LoadFromFile("Service.yaml")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(s, TestStruct) {
		t.Error("Test Failed.")
	}
}

func TestLoadFromJSON(t *testing.T) {
	s, err := LoadFromJSON(TestJSON)
	if err != nil {
		t.Error("Test Failed.")
	}
	if !reflect.DeepEqual(s, TestStruct) {
		t.Error("Test Failed.")
	}
}

func TestToJSON(t *testing.T) {
	_, err := TestStruct.ToJSON()
	if err != nil {
		t.Error("Test Failed.")
	}
}

func TestRequiredKeys(t *testing.T) {
	keys := TestStruct.Parameters.RequiredKeys()
	if keys[0] != "testparam2" {
		t.Error("Test Failed.")
	}
	keys = TestStruct.Flags.RequiredKeys()
	if keys[0] != "test" {
		t.Error("Test Failed.")
	}
}

func TestRequired(t *testing.T) {
	keys := TestStruct.Parameters.Required()
	p := keys[0]
	if p.Key != "testparam2" {
		t.Error("Test Failed.")
	}
	flags := TestStruct.Flags.Required()
	f := flags[0]
	if f.Key != "test" {
		t.Error("Test Failed.")
	}
}

func TestGetParameter(t *testing.T) {
	p, err := TestStruct.Parameters.GetParameter("testparamfake")
	if err == nil {
		t.Error("Expected missing parameter.")
	}
	p, err = TestStruct.Parameters.GetParameter("testparam1")
	if err != nil {
		t.Error("Test Failed.")
	}
	if p.Description != "My first test parameter" {
		t.Error("Test Failed.")
	}
}

func TestGetFlag(t *testing.T) {
	p, err := TestStruct.Flags.GetFlag("testflagfake")
	if err == nil {
		t.Error("Expected missing flag.")
	}
	p, err = TestStruct.Flags.GetFlag("test")
	if err != nil {
		t.Error("Test Failed.")
	}
	if p.Description != "This is a test flag." {
		t.Error("Test Failed.")
	}
}
