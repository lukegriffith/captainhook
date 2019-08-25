package configparser

import (
	"bufio"
	"errors"
	"github.com/lukemgriffith/captainhook"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Endpoints []captainhook.Endpoint `json:"Endpoints"`
	path      string
}

func (c *Config) GetEndpoints() []captainhook.Endpoint {
	return c.Endpoints
}

func (c *Config) Reload() {
	log.Println("loading", c.path)
	c.reload(c.path)
}

func (c *Config) Dump() error {

	d, err := yaml.Marshal(c)

	if err != nil {
		return errors.New("unable to render YAML config from config structure")
	}

	log.Println(string(d))

	return nil
}

func (c *Config) setEndpoint(e []captainhook.Endpoint) {
	c.Endpoints = e
}

func (c *Config) reload(path string) {

	finfo, err := os.Stat(path)
	if err != nil {
		log.Fatal("Unable to determine file information", path)
	}

	switch mode := finfo.Mode(); {
	case mode.IsDir():
		log.Fatal("Unable to load configuration from directory.")
	case mode.IsRegular():
		file, err := os.Open(path)

		if err != nil {
			log.Fatal("Unable to open file", path)
		}

		reader := bufio.NewReader(file)

		data, err := ioutil.ReadAll(reader)

		if err != nil {
			log.Fatal("Unable to read file", path)
		}

		loadedConfig, err := LoadConfig(data)

		if err != nil {
			log.Fatal(err, path)
		}

		c.setEndpoint((loadedConfig.GetEndpoints()))
	}
}

func LoadConfig(data []byte) (*Config, error) {

	c := Config{nil, ""}
	err := yaml.Unmarshal(data, &c)

	log.Println(string(data))

	if err != nil {
		return nil, errors.New("Unable to load config from data.")
	}

	log.Println(c)

	return &c, nil
}

func NewConfig(path string) (*Config, *EndpointService) {

	e := make([]captainhook.Endpoint, 1)

	c := &Config{e, path}

	c.Reload()

	return c, &EndpointService{c}
}
