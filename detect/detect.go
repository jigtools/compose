package detect

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"strings"
)

func Type(filename string) (filetype string, err error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("FAILED to open %s : %s", filename, err)
	}
	if t, err := tryScript(contents); err == nil {
		return t, nil
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

// looks for the leading # (and then at the rest of the file...)
func tryScript(contents []byte) (filetype string, err error) {
	str := string(contents[0:100])
	if strings.HasPrefix(str, "#cloud-config") {
		// TODO: also test that its a vaild yaml and cloud-init (or even if its rancheros / ubuntu etc
		return "cloud-config", nil
	}
	if strings.HasPrefix(str, "#!") {
		return "script", nil
	}
	return "", fmt.Errorf("FAILED to detect a script file")
}
