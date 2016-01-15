// Copyright © 2015 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var serviceFile string
var servicePort string
var serviceHost string

func init() {
	flag.StringVar(&serviceFile, "service-file", "Service.yaml", "Full path to service file.")
	flag.StringVar(&servicePort, "service-port", "", "Port that this service will be operating on. This flag is required")
	flag.StringVar(&serviceHost, "service-host", os.Getenv("SERVICE_HOST"),
		"The hostname this service will be serving from. Overrides SERVICE_HOST environment variable.")
}

// Service definition
type Service struct {

	// Title: Title for service.
	Title string `json:"title" yaml:"Title"`

	// Domain: Domain of the Service many times the github user or organization.
	Domain string `json:"domain" yaml:"Domain"`

	// Version: Version of the Service.
	Version string `json:"version" yaml:"Version"`

	// Type: Category or type of the Service.
	Type string `json:"type" yaml:"Type"`

	// Type: Protocol of the service.
	Protocol string `json:"protocol" yaml:"Protocol"`

	// Private: True if the Service is for internal use only.
	Private bool `json:"private" yaml:"Private"`

	// Requires: An array of Services that are required for this Service,
	// must contain Title, Domain, and Version.
	Requires []Service `json:"requires,omitempty" yaml:"Requires"`

	// Parameters: An array of parameters to call this Service.
	Parameters Parameters `json:"parameters,omitempty" yaml:"Parameters"`

	// Response: A definition of the response structure for this Service.
	Response Response `json:"response" yaml:"Response"`

	// Method: Http method used for this Service.
	Method string `json:"method" yaml:"Method"`

	// Parameters: An array of parameters to call this Service.
	Flags Flags `json:"flags,omitempty" yaml:"Flags"`

	// Port: The Port this service will serve from.  Only applies to local instance.
	Port string `json:"-" yaml:"-"`

	// Host: The hostname from which this service will serve from.  Only applies to local instance.
	Host string `json:"-" yaml:"-"`
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
	s, err := LoadFromFile(serviceFile)
	if err != nil {
		return s, err
	}
	s.Port = servicePort
	s.Host = serviceHost
	return s, err
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

// Load json service definition into current Service
func (s *Service) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// Represents a slice of parameters
type Parameters []Parameter

// Return a slice required parameter keys.
func (ps Parameters) RequiredKeys() []string {
	var required []string
	for _, p := range ps {
		if p.Required {
			required = append(required, p.Key)
		}
	}
	return required
}

// Return a slice required parameters.
func (ps Parameters) Required() []Parameter {
	var required []Parameter
	for _, p := range ps {
		if p.Required {
			required = append(required, p)
		}
	}
	return required
}

// Get Paramater by key.
func (ps Parameters) GetParameter(key string) (Parameter, error) {
	for _, p := range ps {
		if p.Key == key {
			return p, nil
		}
	}
	return Parameter{}, fmt.Errorf("Parameter %v not found", key)
}

// Parameter defines a single parameter for the service to be called.
type Parameter struct {

	// Key: The string key representing the parameter.
	Key string `json:"key" yaml:"Key"`

	// Description: A human readable description of the parameter.
	Description string `json:"description" yaml:"Description"`

	// Required: If the value is required for the API call.
	Required bool `json:"required,omitempty" yaml:"Required"`

	// Type: The type of parameter.  This will be used to identify the
	// location of the parameter in the http.Request.
	Type string `json:"type" yaml:"Type"`

	// Position: A string value representiing a position.  This is relative
	// to the Type.
	Position string `json:"position,omitempty" yaml:"Position"`

	// DataType: A string value that is the key in a map of DataTypes.
	DataType string `json:"dataType" yaml:"DataType"`
}

// Response defines the nature of the response to be returned from this response.
type Response struct {
	// Type is a string identifying the structural definition of the response.
	// The string will reference a value in a map of structural definitions.
	// This is formatting for the response as a whole.
	// An example would be http://github.com/jasonrichardsmith/googlejson
	Type string `json:"type" yaml:"Type"`

	// DataType: A string value that is the key in a map of DataTypes.
	DataType string `json:"dataType" yaml:"DataType"`
}

// Represents a slice of Flags.
type Flags []Flag

// Return a slice required flag keys.
func (fs Flags) RequiredKeys() []string {
	var required []string
	for _, f := range fs {
		if f.Required {
			required = append(required, f.Key)
		}
	}
	return required
}

// Return a slice required flags.
func (fs Flags) Required() []Flag {
	var required []Flag
	for _, f := range fs {
		if f.Required {
			required = append(required, f)
		}
	}
	return required
}

// Get Flag by key.
func (fs Flags) GetFlag(key string) (Flag, error) {
	for _, f := range fs {
		if f.Key == key {
			return f, nil
		}
	}
	return Flag{}, fmt.Errorf("Flag %v not found", key)
}

// Represents a sincle command line flag.
type Flag struct {
	// The flag designation for the command line flag.
	Key string `json:"key" yaml:"Key"`

	// An environment variable that can bes set in lieu of the flag.
	// CLI flag always overrides environment variable.
	Env string `json:"env" yaml:"Env"`

	// Human readable description of the flag.
	Description string `json:"description" yaml:"Description"`

	// Defines flas as required.
	Required bool `json:"required" yaml:"Required"`
}
