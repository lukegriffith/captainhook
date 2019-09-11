package captainhook

import (
	"bytes"
	"testing"
)

func TestRuleExecutes(t *testing.T) {

	var iw bytes.Buffer

	r := Rule{"www.google.com", "{{.test}}"}

	err := r.Execute(&iw, map[string]interface{}{"test": "This is a test"})

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
