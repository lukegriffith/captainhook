package main

import (
	"fmt"
	"github.com/lukemgriffith/captainhook/configparser"
)

func main() {

	c, es := configparser.NewConfig("config.yaml")

	fmt.Println(es.Endpoint("test"))

	c.Reload()

	fmt.Println(es.Endpoint("test"))

}
