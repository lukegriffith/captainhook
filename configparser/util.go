package configparser

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// loads the configuration from the provided path.
func load(path string) ([]byte, error) {

	finfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to determine file information %s", path))
	}

	switch mode := finfo.Mode(); {
	case mode.IsDir():
		return nil, errors.New("unable to load configuration from directory")
	case mode.IsRegular():
		file, err := os.Open(path)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("unable to open file %s", path))
		}

		reader := bufio.NewReader(file)

		data, err := ioutil.ReadAll(reader)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("unable to read file %s", path))
		}

		return data, nil

	}

	return nil, nil
}

// Loads configuration from a byte array performing validation.
func loadConfig(data []byte) (*Config, error) {

	c := Config{nil, ""}
	err := yaml.Unmarshal(data, &c)

	log.Println(string(data))

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func loadSecrets(data []byte) (map[string]string, error) {

	s := make(map[string]string)
	err := yaml.Unmarshal(data, &s)

	log.Println(string(data))

	if err != nil {
		return nil, errors.New("unable to load secrets from path provided")
	}

	return s, nil

}