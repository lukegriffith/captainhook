package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl_string string = `
{
  "bag": [
    "test": {{.test}}
  ]
}
`

type templateStruct struct {
  Name string
  Type string
  ID int
  Bag map[string]int
}



func main() {

  // TODO opposed to passing map directly to template. create custom interface
  // that controls the requests on the object.

  m := map[string]int {
    "test": 123,
    "test1": 345,
  }


	tmpl, err := template.New("rule").Parse(tmpl_string)

	if err != nil {
		log.Fatal("unable to create template", err)
	}

	err = tmpl.Execute(os.Stdout, m)

	if err != nil {
		log.Fatal("failed to write")
	}



}
