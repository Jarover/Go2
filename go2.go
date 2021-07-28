package main

import (
	"fmt"
)

func main() {

	WithWorkers()

	//ch := make(chan int, 3)
	//go GoCounterRoutine(ch)
	//GoWorkerRoutine(1, ch)
	// not run as goroutine because mein() would just end
	//GoWorkerRoutine(2, ch)

}

func WithWorkers() {
	fmt.Println("Start Lesson4")
	var workers = make(chan struct{}, 5)
	var ch = make(chan int, 10)

	ch <- 1
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func(i int) {
			defer func() {
				<-workers
			}()

			val := <-ch
			fmt.Printf("%d -- %d \n", i, val)
			val++
			ch <- val

		}(i)

	}
	val := <-ch
	fmt.Printf("Rrsult %d", val)
}
