package captainhook

import "errors"

type Endpoint struct {
	Name   string   `json:"name"`
	Secret string   `json:"secret"`
	Rules  []Rule   `json:"rules"`
	Sources []Source `json:"sources"`
}

func (e *Endpoint) GetRules() *[]Rule, error {

  if e.Rules == nil {
    return nil, erorrs.New("Endpoint has no associated rules.")
  }
  return e.Rules, nil
}

func (e *Endpoint) GetSources() *[]Source, error {

  if e.Sources == nil {
    return nil, erorrs.New("Endpoint has no associated sources.")
  }
  return e.Sources, nil
}

type EndpointService interface {
	Endpoint(name string) (*Endpoint, error)
	Endpoints() (*[]Endpoint, error)
	CreateEndpoint() error
	DeleteEndpoint() error
}

type Rule struct {
	Destination_url string `json:"destination "`
	Template        string `json:"template"`
	Verify_ssl      string `json:"verify_ssl"`
}

type Source struct {
	SourceType string `json:"sourcetype"`
	Location   string `json:"location"`
}

type SourceType struct {
	Name string `json:"name"`
}
