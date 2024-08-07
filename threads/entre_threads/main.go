package main

import "fmt"

func main() { // Thread 1
	canal := make(chan string)

	// Funcao anonima
	go func() { // Thread 2
		canal <- "valor"
	}()

	msg := <-canal
	fmt.Println(msg)
}

// processo -> Alocar bloco de memoria
// thread -> Acesso ao bloco de memoria
