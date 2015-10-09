package service

import (
	"log"
	"reflect"
	"testing"
)

var TestStruct *Service

func init() {
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
	if s.Title != TestStruct.Title {
		t.Error("Test Failed.")
	}

}
