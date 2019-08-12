package main

import (
	"fmt"
	"github.com/lukemgriffith/captainhook/server"
)

func main() {

	db := server.OpenDB("hello")

	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("unable to access db")
		return
	}

	fmt.Println("db accessed")

}
