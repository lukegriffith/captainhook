package configparser

import (
	"errors"
	"github.com/lukemgriffith/captainhook"
	"github.com/lukemgriffith/captainhook/util"
	"gopkg.in/yaml.v2"
)

// Structure contains the whole configuration when using the config parser
// backend, this is loaded directly from a YAML declaration, and is immutable.
type Config struct {
	Endpoints []captainhook.Endpoint `json:"Endpoints"`
	path      string
  log       *util.Logger
}

// Returns the endpoints of the configuration.
func (c *Config) GetEndpoints() []captainhook.Endpoint {
	return c.Endpoints
}

// Public method to reloads configuration from disk via specified path.
func (c *Config) Reload() error {
	c.log.Println("loading", c.path)
	b, err := load(c.path)

	if err != nil {
		return err
	}

	config, err := loadConfig(b)

	if err != nil {
		return err
	}

	c.setEndpoint(config.GetEndpoints())

	return nil

}

// Dumps YAML Readable configuration to application log for debugging
// purposes.
func (c *Config) Dump() error {

	d, err := yaml.Marshal(c)

	if err != nil {
		return errors.New("unable to render YAML config from config structure")
	}

	c.log.Println(string(d))

	return nil
}

// sets the endpoints that are associated to the configuration.
func (c *Config) setEndpoint(e []captainhook.Endpoint) {
	c.Endpoints = e
}

// Loads configuration from a byte array performing validation.
func loadConfig(data []byte) (*Config, error) {

	c := Config{nil, "", nil}
	err := yaml.Unmarshal(data, &c)

	c.log.Println(string(data))

	if err != nil {
		return nil, err
	}

	return &c, nil
}


// Constructor.
func NewConfig(path string) (*EndpointService, *Config, error) {

	e := make([]captainhook.Endpoint, 1)

	c := &Config{e, path, util.NewDebugLog("config")}

	err := c.Reload()

	if err != nil {
		return nil, nil, err
	}

	return &EndpointService{c}, c, nil
}
