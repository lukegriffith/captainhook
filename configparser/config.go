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
	endpoints *[]captainhook.Endpoint `json:"endpoints"`
	path      string
}

func (c *Config) GetEndpoints() *[]captainhook.Endpoint {
	return c.endpoints
}

func (c *Config) Reload() {
	log.Println("loading", c.path)
	c.reload(c.path)
}

func (c *Config) Dump() {
	log.Println(c.GetEndpoints())
}

func (c *Config) setEndpoint(e *[]captainhook.Endpoint) {

	log.Println("new confg", e)
	c.endpoints = e
}

// TODO: Reload is broken and load a nil structure for endpoints. Needs to be resolved.

func (c *Config) reload(path string) {

	finfo, err := os.Stat(path)
	if err != nil {
		log.Fatal("Unable to determine file information", path)
	}

	switch mode := finfo.Mode(); {
	case mode.IsDir():
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal("Unable to read directory", path)
		}

		for _, f := range files {
			c.reload(f.Name())
		}
	case mode.IsRegular():
		log.Println("2")
		file, err := os.Open(path)

		if err != nil {
			log.Fatal("Unable to open file", path)
		}

		reader := bufio.NewReader(file)

		data, err := ioutil.ReadAll(reader)

		if err != nil {
			log.Fatal("Unable to read file", path)
		}

		log.Println("data", data)

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

	c := &Config{&e, path}

	c.Reload()

	return c, &EndpointService{c}
}
