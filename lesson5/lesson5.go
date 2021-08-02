// Package lesson5 - example functions
package lesson5

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//
//WaitDoneAllGoroutine - запускает n горутин . Дожидается завершения всех запущенных горутин.
//
func WaitDoneAllGoroutine(n int) {
	fmt.Println("Start Lesson 5.1")
	// Создаем и инициализируем семафор
	wg := sync.WaitGroup{}
	wg.Add(n)
	// запускаем n горутин
	for i := 1; i <= n; i++ {
		r := rand.Intn(50000)

		go func(i, d int) {
			time.Sleep(time.Duration(d) * time.Millisecond)
			fmt.Printf("End goroutine %d\n", i)
			wg.Done()

		}(i, r)
	}
	// ждем обнуления семафора
	wg.Wait()
	fmt.Println("All done")
}

//
// UnlocMutexWithDefer - разблокировка мютекса с помощью   defer
//

func UnlocMutexWithDefer(n int) {
	fmt.Println("Start Lesson 5.2")
	var mu sync.Mutex
	var someValue int
	wg := sync.WaitGroup{}
	wg.Add(n)
	// запускаем n горутин
	for i := 1; i <= n; i++ {
		r := rand.Intn(1000)

		go func(i, d int) {
			defer func(i int) {
				fmt.Println("unlock")
				mu.Unlock()
			}(i)
			mu.Lock()
			someValue = i
			time.Sleep(time.Duration(d) * time.Millisecond)
			fmt.Printf("End goroutine %d value = %d ", i, someValue)
			wg.Done()

		}(i, r)
	}
	wg.Wait()
	fmt.Println("All done")

}
