package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl_string string = `
{
  "name": "{{.Name}}",
  "type": "{{.Type}}",
  "ID": "{{.ID}}"
}
`

type templateStruct struct {
	Name string
	Type string
	ID   int
}

func main() {

	t := templateStruct{"Luke", "Human", 231203}

	tmpl, err := template.New("rule").Parse(tmpl_string)

	if err != nil {
		log.Fatal("unable to create template")
	}

	err = tmpl.Execute(os.Stdout, t)

	if err != nil {
		log.Fatal("failed to write")
	}

}
