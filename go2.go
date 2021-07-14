package main

import (
	"fmt"
	"os"

	"github.com/Jarover/Go2/lesson1"
)

func main() {
	filename := os.Args[1]
	content, err := lesson1.ReadMyFile(filename)

	if err != nil {
		fmt.Println("error in ReadMyFile")
	}

	fmt.Println(content)

}
