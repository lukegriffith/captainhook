package configparser

import (
	//"crypto/aes"
	"errors"
	"fmt"
)

type SecretsEndpoint struct {
	data map[string]string
}


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


func(s *SecretsEndpoint) GetTextSecret(name string) (string, error) {

	v, ok := s.data[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("unable to find secret %s", name))
	}

	return v, nil

}