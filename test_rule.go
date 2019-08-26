package captainhook

import (
  "testing"
  "encoding/json"
  "bytes"
)

func Test_rule_execute(t *testing.T) {

  var iw bytes.Buffer

  r := Rule{"www.google.com", "{{.test}}", true}

  err := r.Execute(map[string]interface{}{ "test": 1}, &iw)

  if err != nil {
    t.Fail()
  }

  return
}
