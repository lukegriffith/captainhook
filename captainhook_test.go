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
	r = append(r, Rule{"testURL", "{{.test}}"})
	s = append(s, Source{"Github", "Test"})
	e = Endpoint{"Test", "Secret", r, s}

	rul, err := e.GetRules()
	if err != nil {
		t.Fail()
	}

	if rul[0].Destination != "testURL" {
		t.Fail()
	}
}
