###Service

Is the basic template to define microservices in json and yaml format.

[![GoDoc](https://godoc.org/github.com/duckbunny/service?status.svg)](https://godoc.org/github.com/duckbunny/service)
[![Build Status](https://travis-ci.org/duckbunny/service.svg?branch=master)](https://travis-ci.org/duckbunny/service)
[![Coverage Status](https://coveralls.io/repos/github/duckbunny/service/badge.svg?branch=master)](https://coveralls.io/github/duckbunny/service?branch=master)
# service
--
    import "github.com/duckbunny/service"

Package service is the definition of the microservice.

The Service definition is used to automate bootstrap microservices, automate
communication between separate services and to provide a human readable
definition of a service.


## Usage

```go
var (
	//ErrNoPort when no port has been set for service
	ErrNoPort = errors.New("No port set")
	//ErrNoHost when no host has been set for service
	ErrNoHost = errors.New("No host set")
)
```

#### type Config

```go
type Config struct {
	// Key name for variable
	Key string `json:"key" yaml:"Key"`
	// Required configuration variable
	Required bool `json:"required,omitempty" yaml:"Required"`
	// Description: A human readable description of the parameter.
	Description string `json:"description" yaml:"Description"`
}
```

Config represents one configuration value

#### type Configs

```go
type Configs []Config
```

Configs represents a slice of configs

#### type Flag

```go
type Flag struct {
	// Key is the flag designation for the command line flag.
	Key string `json:"key" yaml:"Key"`

	// Env is an environment variable that can bes set in lieu of the flag.
	// CLI flag always overrides environment variable.
	Env string `json:"env" yaml:"Env"`

	// Description is the human readable description of the flag.
	Description string `json:"description" yaml:"Description"`

	// Required defines flag as required.
	Required bool `json:"required" yaml:"Required"`
}
```

Flag represents a single command line flag.

#### type Flags

```go
type Flags []Flag
```

Flags represents a slice of Flags.

#### func (Flags) GetFlag

```go
func (fs Flags) GetFlag(key string) (Flag, error)
```
GetFlag by key.

#### func (Flags) Required

```go
func (fs Flags) Required() []Flag
```
Required returns a slice required flags.

#### func (Flags) RequiredKeys

```go
func (fs Flags) RequiredKeys() []string
```
RequiredKeys returns a slice required flag keys.

#### type Service

```go
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

	// Configs: An array of configurations this service can use.
	Configs Configs `json:"configs,omitempty" yaml:"Configs"`

	// Parameters: An array of parameters to call this Service.
	Flags Flags `json:"flags,omitempty" yaml:"Flags"`

	// Port: The Port this service will serve from.  Only applies to local instance.
	Port string `json:"-" yaml:"-"`

	// Host: The hostname from which this service will serve from.  Only applies to local instance.
	Host string `json:"-" yaml:"-"`
}
```

Service definition

#### func  LoadFromFile

```go
func LoadFromFile(file string) (*Service, error)
```
LoadFromFile gets a new Service from yaml service definition file.

#### func  LoadFromJSON

```go
func LoadFromJSON(js []byte) (*Service, error)
```
LoadFromJSON loads to new Service from json bytes service definition

#### func  New

```go
func New() *Service
```
New get a new Service.

#### func  This

```go
func This() (*Service, error)
```
This shortcuts to load Service for this application.

#### func (*Service) LoadFromFile

```go
func (s *Service) LoadFromFile(file string) error
```
LoadFromFile loads yaml service definition file into current service.

#### func (*Service) LoadFromJSON

```go
func (s *Service) LoadFromJSON(js []byte) error
```
LoadFromJSON loads service definition into current Service

#### func (*Service) ToJSON

```go
func (s *Service) ToJSON() ([]byte, error)
```
ToJSON converst current service to json
