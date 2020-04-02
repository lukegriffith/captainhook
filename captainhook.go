package captainhook

import "errors"

type (

	// Data structure contains information on an Endpoint, with associated rules
	// and sources.
	Endpoint struct {
		Name    string   `yaml:"name"`
		Secrets []string `yaml:"secrets"`
		Rules   []Rule   `yaml:"rules"`
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

	SecretEngine interface {
		GetTextSecret(name string) (string, error)
		ValidateEndpointConfig(endpoints []Endpoint) error
	}
)

// Obtains the associated rules for an endpoint.
func (e *Endpoint) GetRules() ([]Rule, error) {

	if e.Rules == nil {
		return nil, errors.New("Endpoint has no associated rules.")
	}
	return e.Rules, nil
}
