package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data { // Ler canal
		fmt.Printf("%d - %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // Thread 1
	canal := make(chan int)

	for worker_value := range 25 {
		go worker(worker_value, canal)
	}

	// O canal so recebe 1 valor por vez por padrao
	for value := range 100 {
		canal <- value
	}

}
