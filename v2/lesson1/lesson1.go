package lesson1

import (
	"fmt"
	"os"
	"time"

	_ "gopkg.in/yaml.v3"
)

// Gets the first argument and creates a file
func MyFile() error {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Paniced error = %v\n", r)
		}
	}()

	filename := os.Args[1]

	err := CreateFile(filename)

	if err != nil {
		now := time.Now()
		return WrapLesson1Error(
			fmt.Errorf("Create file %s error: %w ", filename, err),
			now.Unix(),
		)
	}

	fmt.Printf("File %s created!\n", filename)
	return nil
}

// Creates a file and returns an error
func CreateFile(filename string) error {

	file, err := os.Create(filename)

	defer func() {
		err = file.Close()

		if err != nil {
			fmt.Printf("Error close file %s \n", filename)
		}
	}()

	if err != nil {
		return err
	}
	return nil
}
