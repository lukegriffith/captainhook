package configparser

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// loads the configuration from the provided path.
func load(path string) ([]byte, error) {

	finfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to determine file information %s", path))
	}

	switch mode := finfo.Mode(); {
	case mode.IsDir():
		return nil, errors.New("unable to load configuration from directory")
	case mode.IsRegular():
		file, err := os.Open(path)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("unable to open file %s", path))
		}

		reader := bufio.NewReader(file)

		data, err := ioutil.ReadAll(reader)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("unable to read file %s", path))
		}

		return data, nil

	}

	return nil, nil
}


