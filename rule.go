package captainhook

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
)

//TODO: Document
type Rule struct {
	Type 		string `yaml:type`
	Destination string `yaml:destination`
	Arguments   map[string]string `yaml:arguments`
	Function 	func(iw io.Writer, dataMap map[string]interface{}) error
}

func AssignFunction(rule *Rule) {

	switch t := rule.Type; t {
	case "template":
		rule.Function = rule.TemplateFunc

	default:
		rule.Function = rule.NoOp

	}
}

func (rule Rule) GetArg(name string) (string, error) {

	val, ok := rule.Arguments[name]

	if ! ok {
		return "", errors.New(fmt.Sprintf("Unable to find argument %s", name))
	}
	return val, nil
}




func(rule Rule) NoOp (iw io.Writer, dataMap map[string]interface{}) error {
	return nil
}


func(rule Rule) TemplateFunc(iw io.Writer, dataMap map[string]interface{}) error {

	dest, err := rule.GetArg("destination")

	if err != nil {
		return err
	}

	tmplStr, err := rule.GetArg("template")

	if err != nil {
		return err
	}

	tmpl, err := template.New(dest).Parse(tmplStr)

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
