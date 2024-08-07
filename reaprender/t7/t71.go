package main

import (
	"fmt"
	"time"
)

func main() {
	// Uma forma de evitar o lock de memoria
	// é compartilhar ela por meio de uma canal
	// entre as threads

	// O channel fica aberto eternamente esperando uma
	// entrada, então o preempetivo para finalizar o
	// processo não consegue para-lo
	learning := make(chan string)

	go func() {
		learning <- "v1"
	}()

	result := <-learning
	fmt.Println(result)

	// Exemplo do preempetivo
	// O trechos de threads se alternam na execução
	// para devido as regras do scheduler

	outvalue := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		outvalue <- "test"
	}()

	select {
	case vt1 := <-outvalue:
		fmt.Println(vt1)
	default:
		fmt.Println("Default")
	}
}
