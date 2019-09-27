package configparser

import (
	"errors"
	"github.com/lukemgriffith/captainhook"
)

//TODO: Document
type EndpointService struct {
	Config *Config
}

//TODO: Document
func (e *EndpointService) Endpoint(name string) (*captainhook.Endpoint, error) {

	if len(e.Config.GetEndpoints()) == 0 {
		return nil, errors.New("No Endpoints configured")
	}

	for _, endpoint := range e.Config.GetEndpoints() {
		if endpoint.Name == name {
			return &endpoint, nil
		}
	}

	return nil, errors.New("Unable to find endpoint by name")
}

//TODO: Document
func (e *EndpointService) Endpoints() ([]captainhook.Endpoint, error) {
	return e.Config.GetEndpoints(), nil
}

//TODO: Document
func (e *EndpointService) CreateEndpoint() error {
	return errors.New("Unable to create endpoint, defined from static config")
}

//TODO: Document
func (e *EndpointService) DeleteEndpoint() error {
	return errors.New("Unable to delete endpoint, defined from static config")
}
