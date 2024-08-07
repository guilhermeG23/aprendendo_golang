package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go test1("t1")
	go test1("t2")
	time.Sleep(time.Second * 10)

	// Forçar a rodar em um core só
	runtime.GOMAXPROCS(1)

	// Na versão 1.14 adicionaram uma leitura onde
	// se tem loop infinito, ele se da um jeito de sair
	// Isso no scheduler
	go func() {
		for {
		}
	}()
	fmt.Println("Fim")
}

func test1(v1 string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(v1, i)
		time.Sleep(time.Second)
	}
}
