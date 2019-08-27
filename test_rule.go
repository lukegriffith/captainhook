package captainhook

import (
	"bytes"
	"testing"
)

func Test_rule_execute(t *testing.T) {

	var iw bytes.Buffer

	r := Rule{"www.google.com", "{{.test}}", true}

	err := r.Execute(&iw, map[string]interface{}{"test": 1})

	if err != nil {
		t.Fail()
	}

	return
}
