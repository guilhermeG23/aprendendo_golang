// Canal de dados
package main

import "fmt"

func main() {
	canal := make(chan string)

	// Aponta o valor na thread 2 e jogue para o canal
	// So segura um valor por vez, após ser chamado
	// Ele é limpo
	go func() {
		canal <- "Teste"
	}()

	// Printa e limpa o canal
	fmt.Println(<-canal)
}
