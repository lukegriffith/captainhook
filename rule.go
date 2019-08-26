package captainhook

import (
	"bytes"
	"html/template"
	"io"
)

type Rule struct {
	Destination_url string `json:"destination"`
	Template        string `json:"template"`
	Verify_ssl      string `json:"verify_ssl"`
}

func (rule Rule) Execute(iw io.Writer, dataMap map[string]interface{}) error {

	tmpl, err := template.New(rule.Destination_url).Parse(rule.Template)

	if err != nil {
		return err
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, dataMap)

	if err != nil {
		return err
	}

	iw.Write(&tpl)

	return nil
}
