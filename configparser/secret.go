package configparser

import (
	//"crypto/aes"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

// contains implementation of CaptainHook.SecretsEngine.
type SecretsEndpoint struct {
	data map[string]string
}

// Constructor for SecretsEndpoint
func NewSecretEngine(path string, secret string) (*SecretsEndpoint, error) {

	b, err := load(path)

	if err != nil {
		return nil, err
	}

	// TODO: Look into crypto/aes to unencrypt file with secret.

	//secretBytes := []byte(secret)
	//aes.NewCipher(secretBytes)

	sec, err := loadSecrets(b)

	if err != nil {
		return nil, err
	}

	return &SecretsEndpoint{sec}, nil
}


// creates secrets structure from byte array.
func loadSecrets(data []byte) (map[string]string, error) {

	s := make(map[string]string)
	err := yaml.Unmarshal(data, &s)

	log.Println(string(data))

	if err != nil {
		return nil, errors.New("unable to load secrets from path provided")
	}

	return s, nil
}

// interface function, obtains text secret from map.
func(s *SecretsEndpoint) GetTextSecret(name string) (string, error) {

	v, ok := s.data[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("unable to find secret %s", name))
	}

	return v, nil
}