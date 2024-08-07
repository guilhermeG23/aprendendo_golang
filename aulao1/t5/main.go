package main

import (
	"fmt"
	"time"
)

// Canal para compartilhar dados
func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)
	qtdWorker := 1000
	qtdProcess := 10000

	// Simulado de um load balancer
	for x := range qtdWorker {
		go worker(x, canal)
	}

	for i := range qtdProcess {
		canal <- i
	}
}
