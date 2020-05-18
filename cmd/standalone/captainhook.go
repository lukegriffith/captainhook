package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lukemgriffith/captainhook"
	"github.com/lukemgriffith/captainhook/util"
	"github.com/lukemgriffith/captainhook/configparser"
	"github.com/lukemgriffith/captainhook/server"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	var configPath, secretPath, passphrase, port, cryptoFilePath string
	var decrypt, debug bool

	serveSet := flag.NewFlagSet("serve", flag.ExitOnError)
	cryptoSet := flag.NewFlagSet("encrypt", flag.ExitOnError)

	serveSet.StringVar(&configPath, "configPath", "config.yml", "YAML file to configure the service with endpoints.")
	serveSet.StringVar(&secretPath, "secretPath", "", "Encrypted YAML blob containing string map of secrets.")
	serveSet.StringVar(&passphrase, "passphrase", "", "Passphrase for encrypted YAML blob.")
	serveSet.StringVar(&port, "port", ":8081", "TCP port for server to run, default is ':8081'")
  serveSet.BoolVar(&debug, "debug", false, "Should debug messages be printed")

	cryptoSet.StringVar(&cryptoFilePath, "filepath", "","File to perform encryption operation.")
	cryptoSet.StringVar(&passphrase, "passphrase", "", "Passphrase for encrypted YAML blob.")
	cryptoSet.BoolVar(&decrypt, "decrypt", false, "should the file be decrypted")
  cryptoSet.BoolVar(&debug, "debug", false, "Should debug messages be printed")


	if len(os.Args) < 2 {
		fmt.Println("CaptainHook: AWK in the cloud.")
		fmt.Println()
		fmt.Println("captainhook [command] | captainhook [command] -h")
		fmt.Println()
		fmt.Println("Available Commands:")
		fmt.Println("\tserve - Start the CaptainHook application server")
		fmt.Println("\tencrypt - Perform encryption operations")
		fmt.Println("\thelp - Print help for subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		serveSet.Parse(os.Args[2:])
	case "encrypt":
		cryptoSet.Parse(os.Args[2:])
	case "help":
		fmt.Println("serve: Start the CaptainHook application server.")
		serveSet.PrintDefaults()
		fmt.Println()
		fmt.Println("encrypt: Perform encryption operations on a yaml file.")
		cryptoSet.PrintDefaults()
		os.Exit(1)
	default:
		fmt.Println("Invalid subcommand")
		os.Exit(1)
	}

	if serveSet.Parsed() {
    util.SetDebug(debug)
		startServer(configPath, secretPath, passphrase, port)
	}

	if cryptoSet.Parsed() {
    util.SetDebug(debug)
		encryptionCommand(passphrase, cryptoFilePath, decrypt)
	}
}

func encryptionCommand(passphrase string, filepath string, decrypt bool) {

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

func startServer(configPath string, secretPath string, passphrase string, port string) {

	var secEnd captainhook.SecretEngine = createSecretsEngine(secretPath, passphrase)

	svc, config, err := configparser.NewConfig(configPath)

	if err != nil {
		panic(err)
	}

	endpoints, err := svc.Endpoints()

	if err != nil {
		panic(err)
	}

	// Validate that any parsed endpoint that uses secrets,
	// is referencing a secret that exists in the secrets database.
	validateConfig(endpoints, secEnd)

	app := server.New(svc, secEnd)

	hookserver := &http.Server{
		Addr:    port,
		Handler: app,
	}

	go hookserver.ListenAndServe()

	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	exit_channel := make(chan int)
	go func() {
		for {
			s := <-signal_channel
			log.Print("os signal received processing.")
			switch s {
			case syscall.SIGINT:
				config.Reload()
			case syscall.SIGTERM:
				err = hookserver.Shutdown(context.Background())

				if err != nil {
					log.Println(err)
					exit_channel <- 1
				} else {
					exit_channel <- 0
				}
			}
		}
	}()

	code := <-exit_channel
	os.Exit(code)
}




func createSecretsEngine(secretPath string, passphrase string) captainhook.SecretEngine {
	if secretPath != "" && passphrase != "" {
		secEnd, err := configparser.NewSecretEngine(secretPath, passphrase)

		if err != nil {
			panic(err)
		}

		return secEnd
	} else {
		return nil
	}
}



func validateConfig(endpoints []captainhook.Endpoint, secEnd captainhook.SecretEngine) {
	if secEnd != nil {
		err := secEnd.ValidateEndpointConfig(endpoints)

		if err != nil {
			panic(err.Error())
		}
	}
}

