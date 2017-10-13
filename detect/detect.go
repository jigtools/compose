package detect

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Type(filename string) (filetype string, err error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("FAILED to open %s : %s", filename, err)
	}
	if t, err := tryYAML(contents); err == nil {
		return t, nil
	}

	return "", fmt.Errorf("FAILED to detect the file type of %s", filename)
}

func tryYAML(contents []byte) (filetype string, err error) {
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(contents, &m)
	if err != nil {
		return "", fmt.Errorf("FAILED to unmarshal YAML: %s", err)
	}

	return "json", nil
}
