package captainhook

import "errors"

type (

	// Datastructure contains information on an Endpoint, with associated rules
	// and sources.
	Endpoint struct {
		Name    string   `yaml:"name"`
		Secret  string   `yaml:"secret"`
		Rules   []Rule   `yaml:"rules"`
		Sources []Source `yaml:"sources"`
	}

	// Interface provides an extensible way of implementing the EndpointService,
	// this is used in various parts of the application logic to decouple
	// implementations.
	EndpointService interface {
		Endpoint(name string) (*Endpoint, error)
		Endpoints() ([]Endpoint, error)
		CreateEndpoint() error
		DeleteEndpoint() error
	}

	// Contains information on the type of source expected in an endpoint.
	Source struct {
		SourceType string `yaml:"sourcetype"`
		Location   string `yaml:"location"`
	}
)

// Obtains the assocaited rules for an endpoint.
func (e *Endpoint) GetRules() ([]Rule, error) {

	if e.Rules == nil {
		return nil, errors.New("Endpoint has no associated rules.")
	}
	return e.Rules, nil
}

// Obtains the associated sources for an endpoint.
func (e *Endpoint) GetSources() ([]Source, error) {

	if e.Sources == nil {
		return nil, errors.New("Endpoint has no associated sources.")
	}
	return e.Sources, nil
}
