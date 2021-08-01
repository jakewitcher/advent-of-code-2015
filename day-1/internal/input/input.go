package input

import "io/ioutil"

func Extract(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
