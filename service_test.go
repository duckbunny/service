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
		Title:   "service",
		Domain:  "duckbunny",
		Version: "0.1",
		Type:    "test",
		Private: false,
		Method:  "POST",
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
	}
}

func TestThis(t *testing.T) {
	s, err := This()
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(s, TestStruct) {
		t.Error("Test Failed.")
	}
}

func TestLoadFromFile(t *testing.T) {
	s, err := LoadFromFile("Service.yaml")
	if err != nil {
		log.Fatal(err)
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
}

func TestRequired(t *testing.T) {
	keys := TestStruct.Parameters.Required()
	p := keys[0]
	if p.Key != "testparam2" {
		t.Error("Test Failed.")
	}
}

func TestGetParameter(t *testing.T) {
	p, err := TestStruct.Parameters.GetParameter("testparam1")
	if err != nil {
		t.Error("Test Failed.")
	}
	if p.Description != "My first test parameter" {
		t.Error("Test Failed.")
	}
}
