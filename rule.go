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
	Secrets		[]string 		  `yaml:secrets`
	function    func(iw io.Writer, dataMap map[string]interface{}, r *Rule) error
}

// Executes the function poitner associated to the rule.
func (r *Rule) Execute(iw io.Writer, dataMap map[string]interface{}) error {
	if r.function == nil {
		return errors.New("No function pointer assigned.")
	}

	return r.function(iw, dataMap, r)
}

// Sets the function pointer of the rule.
func (r *Rule) SetFunction(f func(iw io.Writer, dataMap map[string]interface{}, r *Rule) error) error {

	if r.function != nil {
		return errors.New("Function already exists.")
	}

	r.function = f

	return nil
}

// TODO make this mapping more configurable./find out if there's a better way to do this.
// Maps the function to the rule type.
func AssignFunction(rule *Rule) {

	switch t := rule.Type; t {
	case "template":
		rule.SetFunction(TemplateFunc)

	default:
		rule.SetFunction(NoOp)

	}
}

// Checks the rule arguments map for value and returns it or error.
func (rule Rule) GetArg(name string) (string, error) {

	val, ok := rule.Arguments[name]

	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find argument %s", name))
	}
	return val, nil
}


// RULE FUNCTIONS.
// Rules are passed a io.Writer, what wil be used in the POST request to downstream services.
// dataMaps are arguments of the executing rule, passed in from the configuration or input.
// secrets map is passed in.


// NoOp, do nothing.
func NoOp(iw io.Writer, dataMap map[string]interface{}, rule *Rule) error {
	return nil
}


// Uses go templating to create a new json string from the input recieved.
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
