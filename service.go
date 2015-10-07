// Copyright © 2015 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package service

import (
	"flag"
	"github.com/mongolar/service/parameter"
	"log"
	"os"
)

var serviceFile string
var DefaultService *Service

func init() {
	DefaultService = New()
	flag.StringVar(&serviceFile, "service", "Service.yaml", "Full path to service file.")
}

// Service definition
type Service struct {
	Title      string               `json:"Title"`
	Domain     string               `json:"Domain"`
	Version    string               `json:"Version"`
	Type       string               `json:"Type"`
	Private    bool                 `json:"Private"`
	Requires   []Service            `json:"Requires,omitempty"`
	Parameters parameter.Parameters `json:"Parameters"`
	Response   Response             `json:"Response"`
	Method     string               `json:"Method"`
}

// Get a new Service and set the default Handler to the DefaultServerMux
func New() *Service {
	service := new(Service)
	return service
}

// Shortcut to load Default Config Service
func LoadConfig() {
	DefaultService.LoadConfigFile()
}

// Load config from flag setting
func (s *Service) LoadConfigFile() {
	if !flag.Parsed() {
		flag.Parse()
	}
	_, err := os.Stat(serviceFile)
	if err == nil {
		filename := len(file) - len(".yaml")
		v := viper.New()
		v.SetConfigName(file[0:filename])
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
		err = v.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	log.Fatal(err)
}

type Parameters []Parameter

// Return all required parameters.
func (ps Parameters) GetRequired() []string {
	required := make([]string, 0)
	for _, p := range ps {
		if p.Required {
			required = append(required, p.Key)
		}
	}
	return required
}

func (ps Parameters) Get(key string) (*Parameter, error) {
	for _, p := range ps {
		if p.Key == key {
			return &p, nil
		}
	}
	return new(Parameter), fmt.Errorf("Parameter %v not found", key)
}

/*
Parameter defines a single parameter for the service to be called.

Title: Is a human readable title for the parameter, it will be used as a
value key for form values.

Type: The type is one of the following default types:
        "form" = normal form post value
        "url" = part of the url string (requires position)
        "json" = submitted as json
        "query" = as a query parameter in the url

Additional values can be added with the AddType function.

Position: is only relevant to url types, determines position in url.

Required: Required value for service.


*/
type Parameter struct {
	Key         string
	Description string
	Type        string
	Required    bool
	DataType    string
	Position    string
}
