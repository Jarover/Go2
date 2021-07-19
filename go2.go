package main

import (
	"fmt"

	"github.com/Jarover/Go2/lesson1"
)

func main() {
	err := lesson1.MyFile()
	if err != nil {
		fmt.Println("Error", err)
	}

}
