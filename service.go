// Copyright © 2015 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package service

import (
	"encoding/json"
	"flag"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var serviceFile string

func init() {
	flag.StringVar(&serviceFile, "service-file", "Service.yaml", "Full path to service file.")
}

// Service definition
type Service struct {

	// Title: Title for service.
	Title string `json:"title"`

	// Domain: Domain of the Service many times the github user or organization.
	Domain string `json:"domain"`

	// Version: Version of the Service.
	Version string `json:"version"`

	// Type: Category or type of the Service.
	Type string `json:"type"`

	// Private: True if the Service is for internal use only.
	Private bool `json:"private"`

	// Requires: An array of Services that are required for this Service,
	// must contain Title, Domain, and Version.
	Requires []Service `json:"requires,omitempty"`

	// Parameters: An array of parameters to call this Service.
	Parameters Parameters `json:"parameters,omitempty"`

	// Response: A definition of the response structure for this Service.
	Response Response `json:"response"`

	// Method: Http method used for this Service.
	Method string `json:"method"`
}

// Get a new Service.
func New() *Service {
	service := new(Service)
	return service
}

// Shortcut to load Service for this application.
func This() (*Service, error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	return LoadFromFile(serviceFile)
}

// Shortcut to get new Service from yaml service definition file.
func LoadFromFile(file string) (*Service, error) {
	s := New()
	err := s.LoadFromFile(file)
	return s, err
}

// Load yaml service definition file into current service.
func (s *Service) LoadFromFile(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		filebytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(filebytes, s)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

// Shortcut to new Service from json bytes service definition
func LoadFromJSON(js []byte) (*Service, error) {
	s := New()
	err := s.LoadFromJSON(js)
	return s, err
}

// Load json service definition into current Service
func (s *Service) LoadFromJSON(js []byte) error {
	return json.Unmarshal(js, s)
}

// Represents a slice of parameters
type Parameters []Parameter

// Return a slice required parameter keys.
func (ps Parameters) RequiredKeys() []string {
	required := make([]string, 0)
	for _, p := range ps {
		if p.Required {
			required = append(required, p.Key)
		}
	}
	return required
}

// Return a slice required parameters.
func (ps Parameters) Required() []Parameter {
	required := make([]Parameter, 0)
	for _, p := range ps {
		if p.Required {
			required = append(required, p)
		}
	}
	return required
}

// Get Paramater by key.
func (ps Parameters) GetParameter(key string) (*Parameter, error) {
	for _, p := range ps {
		if p.Key == key {
			return &p, nil
		}
	}
	return new(Parameter), fmt.Errorf("Parameter %v not found", key)
}

// Parameter defines a single parameter for the service to be called.
type Parameter struct {

	// Key: The string key representing the parameter.
	Key string `json:"key"`

	// Description: A human readable description of the parameter.
	Description string `json:"description"`

	// Required: If the value is required for the API call.
	Required bool `json:"required,omitempty"`

	// Type: The type of parameter.  This will be used to identify the
	// location of the parameter in the http.Request.
	Type string `json:"type"`

	// Position: A string value representiing a position.  This is relative
	// to the Type.
	Position string `json:"position,omitempty"`

	// DataType: A string value that is the key in a map of DataTypes.
	DataType string `json:"dataType"`
}

// Response defines the nature of the response to be returned from this response.
type Response struct {
	// Type is a string identifying the structural definition of the response.
	// The string will reference a value in a map of structural definitions.
	// This is formatting for the response as a whole.
	// An example would be http://github.com/jasonrichardsmith/googlejson
	Type string `json:"type"`

	// DataType: A string value that is the key in a map of DataTypes.
	DataType string `json:"dataType"`
}
