package lesson1

import (
	"bytes"
	"os"
)

func ReadMyFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	return contents, nil
}
