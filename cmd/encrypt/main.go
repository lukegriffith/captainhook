package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/lukemgriffith/captainhook/configparser"
)

func main() {

	var passphrase string
	var filepath string
	var decrypt bool

	flag.StringVar(&passphrase, "passphrase", "", "passphrase to encrypt data with")
	flag.StringVar(&filepath, "filepath", "", "path to file to encrypt")
	flag.BoolVar(&decrypt, "decrypt", false, "should the file be decrypted")

	flag.Parse()

	if passphrase == "" || filepath == "" {
		flag.Usage()
		os.Exit(0)
	}

	key := configparser.NewSymmetricKey(passphrase)

	file, err := os.Open(filepath)

	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err.Error())
	}

	if decrypt {
		text := key.Decrypt(data)

		err = ioutil.WriteFile(filepath, text, 0777)

		if err != nil {
			panic(err.Error())
		}

	} else {
		ciphertext := key.Encrypt(data)

		err = ioutil.WriteFile(filepath, ciphertext, 0777)

		if err != nil {
			panic(err.Error())
		}
	}

}
