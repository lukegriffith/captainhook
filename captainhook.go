package captainhook

import "errors"

type (
	Endpoint struct {
		Name    string   `json:"name"`
		Secret  string   `json:"secret"`
		Rules   []Rule   `json:"rules"`
		Sources []Source `json:"sources"`
	}

	EndpointService interface {
		Endpoint(name string) (*Endpoint, error)
		Endpoints() ([]Endpoint, error)
		CreateEndpoint() error
		DeleteEndpoint() error
	}

	Source struct {
		SourceType string `json:"sourcetype"`
		Location   string `json:"location"`
	}

	SourceType struct {
		Name string `json:"name"`
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
