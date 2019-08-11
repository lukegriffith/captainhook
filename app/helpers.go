package app

import (
	"errors"
	"net/url"
	"strings"
)

func getID(u *url.URL) (string, error) {
	path := u.Path
	parts := strings.Split(path, "/")

	if len(parts) == 1 {
		return "", errors.New("Unable to parse ID from URL")
	}
	return parts[len(parts)-1], nil
}
