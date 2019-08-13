package main

import (
	"fmt"
	"github.com/lukemgriffith/captainhook/configparser"
)

var data = `
endpoints:
  - name: test
    secret: test
`




func main() {


  _, es := configparser.NewConfig(data)


  fmt.Println(es.Endpoint("test1"))

}
