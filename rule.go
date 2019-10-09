package captainhook

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
)

// Contains rule state, is assigned a function based on type.
// Type mapped by AssignFunction.
type Rule struct {
	Type        string            `yaml:type`
	Destination string            `yaml:destination`
	Arguments   map[string]string `yaml:arguments`
	Function    func(iw io.Writer, dataMap map[string]interface{}, r *Rule) error
}

// TODO make this mapping more configurable.
// TODO find out if theres a better way to do this.
// Maps the function to the rule type.
func AssignFunction(rule *Rule) {

	switch t := rule.Type; t {
	case "template":
		rule.Function = TemplateFunc

	default:
		rule.Function = NoOp

	}
}

func (rule Rule) GetArg(name string) (string, error) {

	val, ok := rule.Arguments[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find argument %s", name))
	}
	return val, nil
}

func NoOp(iw io.Writer, dataMap map[string]interface{}, rule *Rule) error {
	return nil
}

func TemplateFunc(iw io.Writer, dataMap map[string]interface{}, rule *Rule) error {

	tmplStr, err := rule.GetArg("template")

	if err != nil {
		return err
	}

	tmpl, err := template.New("tmpl").Parse(tmplStr)

	if err != nil {
		return err
	}

	buf := make([]byte, 0, 1)
	var tpl *bytes.Buffer = bytes.NewBuffer(buf)

	err = tmpl.Execute(tpl, dataMap)

	if err != nil {
		return err
	}

	_, err = iw.Write(tpl.Bytes())

	if err != nil {
		return err
	}

	return nil
}
