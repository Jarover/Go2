package lesson1

import (
	"bytes"
	"fmt"
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

func CreateFile(filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()

		// Обратите внимание: при закрытии файла также возможна ошибка
		if err != nil {
			fmt.Printf("Error close file %s", filename)
		}
	}()
	return nil
}
