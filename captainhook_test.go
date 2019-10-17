 package captainhook

import (
	"testing"
)

// Ensures the endpoint services constructs correctly and rules are obtained.
func TestEndpoint(t *testing.T) {

	var (
		r []Rule
		s []Source
		e Endpoint
	)

	args := make(map[string]string)
	args["template"] = "{{.test}}"

	r = append(r, Rule{"template", "testURL", args, nil, TemplateFunc})
	s = append(s, Source{"Github", "Test"})

	secrets := []string{"test", "test1"}

	e = Endpoint{"Test", secrets, r, s}

	rul, err := e.GetRules()
	if err != nil {
		t.Fail()
	}

	if rul[0].Destination != "testURL" {
		t.Fail()
	}
}
