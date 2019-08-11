package app

import (
	"net/url"
	"testing"
)

func Test_getID_parses(t *testing.T) {

	u, err := url.Parse("https://test/123")
	if err != nil {
		t.Fail()
	}
	id, err := getID(u)
	if id != "123" || err != nil {
		t.Fail()
	}
}

func Test_getID_fails(t *testing.T) {

	u, err := url.Parse("https://test")
	if err != nil {
		t.Fail()
	}
	_, err = getID(u)
	if err == nil {
		t.Fail()
	}
}
