package captainhook

import (
	"bytes"
	"html/template"
	"io"
)

type Rule struct {
	Destination_url string `json:"destination"`
	Template        string `json:"template"`
	Verify_ssl      bool `json:"verify_ssl"`
}

func (rule Rule) Execute(iw io.Writer, dataMap map[string]interface{}) error {

	tmpl, err := template.New(rule.Destination_url).Parse(rule.Template)

	if err != nil {
		return err
	}


  buf := make([]byte, 0, 1)
	var tpl *bytes.Buffer = bytes.NewBuffer(buf)

	err = tmpl.Execute(tpl, dataMap)

	if err != nil {
		return err
	}

	iw.Write(tpl.Bytes())

	return nil
}
