package captainhook

import (
	"bytes"
	"testing"
)

//TODO: Document
func TestRuleExecutes(t *testing.T) {

	var iw bytes.Buffer

	args := make(map[string]string)
	args["template"] = "{{.test}}"

	r := Rule{"template", "www.google.com", args, nil, false, TemplateFunc}

	err := r.function(&iw, map[string]interface{}{"test": "This is a test"}, &r)

	if err != nil {
		t.Fail()
	}

	tmplStr := iw.String()

	t.Log("template:", tmplStr)

	if tmplStr != "This is a test" {
		t.Fail()
	}

	return
}
