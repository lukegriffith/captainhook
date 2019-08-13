package configparser

import (
  "github.com/lukemgriffith/captainhook"
	"gopkg.in/yaml.v2"
  "errors"
)


type Config struct {
  Endpoints []captainhook.Endpoint
}

func LoadConfig(data string) (*Config, error) {

  c := Config{}
  err := yaml.Unmarshal([]byte(data), &c)

  if err != nil {
    return nil, errors.New("Unable to load config from data.")
  }

  return &c, nil
}

func NewConfig(data string) (*Config, *EndpointService) {

  c, err := LoadConfig(data)

  if err != nil {
    return nil, nil
  }

  return c, &EndpointService{c}
}
