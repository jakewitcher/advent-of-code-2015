package input

import (
	"io/ioutil"
	"strings"
)

func Extract(path string) ([]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	values := strings.Split(string(file), "\r\n")

	return values, nil
}
