###Service

Is the basic template to define microservices in json and yaml format.

[![GoDoc](https://godoc.org/github.com/duckbunny/service?status.svg)](https://godoc.org/github.com/duckbunny/service)


# service
--
    import "github.com/duckbunny/service"

Service is the definition of the microservice.

The Service definition is used to automate bootstrap microservices, automate
communication between separate services and to provide a human readable
definition of a service.


## Usage

#### type Flag

```go
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
```

Represents a sincle command line flag.

#### type Flags

```go
type Flags []Flag
```

Represents a slice of Flags.

#### func (Flags) GetFlag

```go
func (fs Flags) GetFlag(key string) (Flag, error)
```
Get Flag by key.

#### func (Flags) Required

```go
func (fs Flags) Required() []Flag
```
Return a slice required flags.

#### func (Flags) RequiredKeys

```go
func (fs Flags) RequiredKeys() []string
```
Return a slice required flag keys.

#### type Parameter

```go
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
```

Parameter defines a single parameter for the service to be called.

#### type Parameters

```go
type Parameters []Parameter
```

Represents a slice of parameters

#### func (Parameters) GetParameter

```go
func (ps Parameters) GetParameter(key string) (Parameter, error)
```
Get Paramater by key.

#### func (Parameters) Required

```go
func (ps Parameters) Required() []Parameter
```
Return a slice required parameters.

#### func (Parameters) RequiredKeys

```go
func (ps Parameters) RequiredKeys() []string
```
Return a slice required parameter keys.

#### type Response

```go
type Response struct {
	// Type is a string identifying the structural definition of the response.
	// The string will reference a value in a map of structural definitions.
	// This is formatting for the response as a whole.
	// An example would be http://github.com/jasonrichardsmith/googlejson
	Type string `json:"type" yaml:"Type"`

	// DataType: A string value that is the key in a map of DataTypes.
	DataType string `json:"dataType" yaml:"DataType"`
}
```

Response defines the nature of the response to be returned from this response.

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
```

Service definition

#### func  LoadFromFile

```go
func LoadFromFile(file string) (*Service, error)
```
Shortcut to get new Service from yaml service definition file.

#### func  LoadFromJSON

```go
func LoadFromJSON(js []byte) (*Service, error)
```
Shortcut to new Service from json bytes service definition

#### func  New

```go
func New() *Service
```
Get a new Service.

#### func  This

```go
func This() (*Service, error)
```
Shortcut to load Service for this application.

#### func (*Service) LoadFromFile

```go
func (s *Service) LoadFromFile(file string) error
```
Load yaml service definition file into current service.

#### func (*Service) LoadFromJSON

```go
func (s *Service) LoadFromJSON(js []byte) error
```
Load json service definition into current Service

#### func (*Service) ToJSON

```go
func (s *Service) ToJSON() ([]byte, error)
```
Load json service definition into current Service
