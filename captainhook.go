package captainhook

type Endpoint struct {
	Name   string   `json:"name"`
	Secret string   `json:"secret"`
	Rules  []Rule   `json:"rules"`
	Source []Source `json:"sources"`
}

type EndpointService interface {
  Endpoint(name string) (*Endpoint, error)
  Endpoints() ([]*Endpoint, error)
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

