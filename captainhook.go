package captainhook

import "errors"

type (

	//TODO: Document
	Endpoint struct {
		Name    string   `yaml:"name"`
		Secret  string   `yaml:"secret"`
		Rules   []Rule   `yaml:"rules"`
		Sources []Source `yaml:"sources"`
	}

	//TODO: Document
	EndpointService interface {
		Endpoint(name string) (*Endpoint, error)
		Endpoints() ([]Endpoint, error)
		CreateEndpoint() error
		DeleteEndpoint() error
	}

	//TODO: Document
	Source struct {
		SourceType string `yaml:"sourcetype"`
		Location   string `yaml:"location"`
	}

	//TODO: Document
	SourceType struct {
		Name string `yaml:"name"`
	}
)

//TODO: Document
func (e *Endpoint) GetRules() ([]Rule, error) {

	if e.Rules == nil {
		return nil, errors.New("Endpoint has no associated rules.")
	}
	return e.Rules, nil
}

//TODO: Document
func (e *Endpoint) GetSources() ([]Source, error) {

	if e.Sources == nil {
		return nil, errors.New("Endpoint has no associated sources.")
	}
	return e.Sources, nil
}
