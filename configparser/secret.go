package configparser

import (
	//"crypto/aes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/lukemgriffith/captainhook"
	"io"
	"gopkg.in/yaml.v2"
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

	sec, err := loadSecrets(b, secret)

	if err != nil {
		return nil, err
	}

	return &SecretsEndpoint{sec}, nil
}

// creates secrets structure from byte array.
func loadSecrets(cipherText []byte, passphrase string) (map[string]string, error) {

	k := NewSymmetricKey(passphrase)
	data := k.Decrypt(cipherText)
	s := make(map[string]string)
	err := yaml.Unmarshal(data, &s)

	if err != nil {
		return nil, errors.New("unable to load secrets from path provided")
	}

	return s, nil
}

// Obtains text secret from map.
func (s *SecretsEndpoint) GetTextSecret(name string) (string, error) {

	v, ok := s.data[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("unable to find secret %s", name))
	}

	return v, nil
}

// Validating all configured secrets are available.
func (s *SecretsEndpoint) ValidateEndpointConfig(endpoints []captainhook.Endpoint) error {

	var errorsFound = false
	var missingSecrets []string

	for _, end := range endpoints {
		for _, name := range end.Secrets {
			_, err := s.GetTextSecret(name)
			if err != nil {
				errorsFound = true
				missingSecrets = append(missingSecrets, name)

			}
		}
	}

	if errorsFound {
		return errors.New(fmt.Sprintf("unable to find secret %q", missingSecrets))
	}

	return nil
}

// Struct contains hashed key for encryption / decryption
type SymmetricKey struct {
	key string
}

// NewSymmetricKey creates an object to deal with aes encryption from provided passphrase
func NewSymmetricKey(key string) *SymmetricKey {
	hasher := md5.New()
	hasher.Write([]byte(key))

	return &SymmetricKey{hex.EncodeToString(hasher.Sum(nil))}
}

// Decrypts provided data using key.
func (k *SymmetricKey) Decrypt(data []byte) []byte {

	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

// Encrypts provided data using key.
func (k *SymmetricKey) Encrypt(data []byte) []byte {
	block, _ := aes.NewCipher([]byte(k.key))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}
