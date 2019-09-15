package captainhook

import "errors"

type (
	Endpoint struct {
		Name    string   `yaml:"name"`
		Secret  string   `yaml:"secret"`
		Rules   []Rule   `yaml:"rules"`
		Sources []Source `yaml:"sources"`
	}

	EndpointService interface {
		Endpoint(name string) (*Endpoint, error)
		Endpoints() ([]Endpoint, error)
		CreateEndpoint() error
		DeleteEndpoint() error
	}

	Source struct {
		SourceType string `yaml:"sourcetype"`
		Location   string `yaml:"location"`
	}

	SourceType struct {
		Name string `yaml:"name"`
	}
)

func (e *Endpoint) GetRules() ([]Rule, error) {

	if e.Rules == nil {
		return nil, errors.New("Endpoint has no associated rules.")
	}
	return e.Rules, nil
}

func (e *Endpoint) GetSources() ([]Source, error) {

	if e.Sources == nil {
		return nil, errors.New("Endpoint has no associated sources.")
	}
	return e.Sources, nil
}
