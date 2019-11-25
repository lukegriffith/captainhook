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
	"io"
	"log"

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
func loadSecrets(ciphertext []byte, passphrase string) (map[string]string, error) {

	k := NewAsymmetricKey(passphrase)
	data := k.Decrypt(ciphertext)
	s := make(map[string]string)
	err := yaml.Unmarshal(data, &s)

	log.Println(string(data))

	if err != nil {
		return nil, errors.New("unable to load secrets from path provided")
	}

	return s, nil
}

// interface function, obtains text secret from map.
func (s *SecretsEndpoint) GetTextSecret(name string) (string, error) {

	v, ok := s.data[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("unable to find secret %s", name))
	}

	return v, nil
}

// Utility functions

type asymmetricKey struct {
	key string
}

// NewAsymmetricKey gcreates an object to deal with aes encryption from provided passphrase
func NewAsymmetricKey(key string) *asymmetricKey {
	hasher := md5.New()
	hasher.Write([]byte(key))

	return &asymmetricKey{hex.EncodeToString(hasher.Sum(nil))}
}

func (k *asymmetricKey) Decrypt(data []byte) []byte {

	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

func (k *asymmetricKey) Encrypt(data []byte) []byte {
	block, _ := aes.NewCipher([]byte(k.key))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
