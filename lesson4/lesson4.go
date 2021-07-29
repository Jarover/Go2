// Package lesson4 - example functions
package lesson4

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//
//WithWorkers - запускает 1000 горутин в пуле воркеров. Горутины получают через канал число, инкрементируют его и возвращают в канал.
//
func WithWorkers() {
	fmt.Println("Start Lesson4.1")
	// канал для воркеров
	var workers = make(chan struct{}, 5)
	// Канал для обмена результатами работы
	var ch = make(chan int, 1)

	ch <- 1
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func(i int) {
			defer func() {
				<-workers
			}()
			// получаем число из канала
			val := <-ch
			val++
			// отправляем инкрементированное число в канал
			ch <- val
			fmt.Printf("Goroutine %d -- Result %d \n", i, val)

		}(i)

	}
	// получаем финальное число и выводим результат
	fmt.Printf("Result %d", <-ch)
}

//WithSigterm - Запускает 1000 горутин со случайным временем исполнения
//и ждет получения сигнала SIGTERM для завершения работы не позднее 1 сек.
func WithSigterm() {
	fmt.Println("Start Lesson4.2")
	ctx, cancelFunc := context.WithCancel(context.Background())
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	/*
		defer func() {
			fmt.Println("Отмена контекста")
			cancelFunc()
		}()
	*/

	// запускаем 1000 горутин
	for i := 1; i <= 1000; i++ {

		go func(ctx context.Context, i int) {

			timeExit := time.Duration(rand.Intn(10000)) * time.Millisecond

			select {
			case <-ctx.Done():
				// Если контекст истекает, выбирается этот случай
				fmt.Printf("Done goroutine %d by cansel  \n", i)

			case <-time.After(timeExit):
				// Этот вариант выбирается, когда работа завершается до отмены контекста
				fmt.Printf("Done goroutine %d by timeout  \n", i)
			}

		}(ctx, i)

	}
	// Создаем  канал
	termCh := make(chan os.Signal)
	signal.Notify(termCh, syscall.SIGINT, syscall.SIGTERM)

	<-termCh // Ждем сигнал SIGTERM
	fmt.Println("***\nShutdown signal received\n***")
	// ждем секунду
	cancelFunc()
	//time.Sleep(10 * time.Second)
	fmt.Println("***\nWait 1 second and stoped\n***")
}

// sleepFunc - функция засыпающая на случайное время
func sleepFunc(i int) {
	r := rand.Intn(i)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
