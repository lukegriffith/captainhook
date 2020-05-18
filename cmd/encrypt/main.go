package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/lukemgriffith/captainhook/configparser"
)

func main() {

	var passphrase, filepath, cryptoFilePath string
	var decrypt bool

	serveSet := flag.NewFlagSet("serve", flag.ExitOnError)
	cryptoSet := flag.NewFlagSet("encrypt", flag.ExitOnError)

	serveSet.StringVar(&passphrase, "passphrase", "", "passphrase to encrypt data with")
	serveSet.StringVar(&filepath, "filepath", "", "path to file to encrypt")
	serveSet.BoolVar(&decrypt, "decrypt", false, "should the file be decrypted")

	cryptoSet.StringVar(&cryptoFilePath, "filepath", "","File to perform encryption operation.")



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
