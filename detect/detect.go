package detect

import "fmt"

func Type(filename string) (filetype string, err error) {
	return "", fmt.Errorf("FAILED to detect the file type of %s", filename)
}
