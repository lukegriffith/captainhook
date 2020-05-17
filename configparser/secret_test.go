package configparser

import (
	"github.com/lukemgriffith/captainhook"
	"io/ioutil"
	"os"
	"testing"
)


const (
	testData = `
testKey: 123
`
	passphrase = "mysecretkey"
)
func TestEncryptionDecryption(t *testing.T) {

	k := NewSymmetricKey(passphrase)
	plaintext := []byte("This is a test string")
	ciphertext := k.Encrypt(plaintext)

	unencryptedPlain := k.Decrypt(ciphertext)

	if string(unencryptedPlain) != string(plaintext) {
		t.Fail()
	}

}


func TestSecretsMap(t *testing.T) {


	k := NewSymmetricKey(passphrase)
	plaindata := []byte(testData)
	ciphertext := k.Encrypt(plaindata)

	secrets, err := loadSecrets(ciphertext, passphrase)

	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	secEnd := &SecretsEndpoint{secrets}

	testKey, err := secEnd.GetTextSecret("testKey")

	if  testKey != "123" {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestLoadFromFile(t *testing.T) {

	file, err := ioutil.TempFile("", "tempfile")
	if err != nil{
		t.Log(err.Error())
		t.Fail()
	}

	defer file.Close()
	defer os.Remove(file.Name())

	fpath := file.Name()

	k := NewSymmetricKey(passphrase)
	plaindata := []byte(testData)
	ciphertext := k.Encrypt(plaindata)

	_, err = file.Write([]byte(ciphertext))

	if err != nil{
		t.Log(err.Error())
		t.FailNow()
	}

	secEnd, err := NewSecretEngine(fpath, passphrase)

	if err != nil{
		t.Log(err.Error())
		t.FailNow()
	}

	testKey, err := secEnd.GetTextSecret("testKey")

	if err != nil{
		t.Log(err.Error())
		t.FailNow()
	}

	if  testKey != "123" {
		t.Log(err.Error())
		t.FailNow()
	}
}


func TestValidation(t *testing.T) {


	k := NewSymmetricKey(passphrase)
	plaindata := []byte(testData)
	ciphertext := k.Encrypt(plaindata)

	secrets, err := loadSecrets(ciphertext, passphrase)

	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	secEnd := &SecretsEndpoint{secrets}

	validEndpoint := captainhook.Endpoint{"TestEnd", []string{"testKey"}, nil}
	invalidEndpoint := captainhook.Endpoint{"TestEnd", []string{"testKey", "fakeKey"}, nil}

	var endpoints []captainhook.Endpoint

	endpoints = append(endpoints, validEndpoint)

	err = secEnd.ValidateEndpointConfig(endpoints)

	if err != nil {
		t.Log("validEndpoint threw validation error on SecEnd")
		t.FailNow()
	}

	endpoints = append(endpoints, invalidEndpoint)

	err = secEnd.ValidateEndpointConfig(endpoints)

	if err == nil {
		t.Log("invalid endpoint did not throw a validation error")
		t.FailNow()
	}


}