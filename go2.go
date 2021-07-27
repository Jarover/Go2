package main

import (
	"fmt"
	"time"

	"github.com/Jarover/Go2/lesson4"
)

func main() {
	fmt.Println("Start Lesson4")

	counter := new(lesson4.Counter)

	for i := 1; i <= 1000; i++ {
		go counter.Incriment()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(counter.Result())
}
