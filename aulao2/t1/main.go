package main

import "fmt"

func main() {
	println("teste")
	fmt.Println("teste")

	// Atribuicao
	var t1 string
	t1 = "teste"
	fmt.Println(t1)
	fmt.Printf("T %s\n", t1)

	t2 := "teste"
	fmt.Println(t2)

	// Funcao com retorno
	fmt.Println(soma(1, 2))
	fmt.Println(soma1(1, 2))
}

func soma(v1 int, v2 int) int {
	return v1 + v2
}

func soma1(v1 int, v2 int) (int, bool) {
	return v1 + v2, true
}
