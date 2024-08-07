package main

import (
	"fmt"

	"t2/guilherme"
)

func main() {
	// Declaração completa
	// Aqui vc declara qual é exatamente o tipo da variavel
	var t1 string
	t1 = "t1"
	fmt.Println(t1)

	// Declaração simples
	// Aqui ele adivinha com base no valor
	t2 := "t2"
	fmt.Println(t2)

	// Nao pode fazer um segundo t2 :=
	// t2 := "teste" // O motivo é que o : server para fazer
	// a declaração do tipo junto, dai para frente
	// vc tem que só usar o = para alterar o valor

	t3 := 1
	t4 := true
	t5 := 1.1
	// Valores multiline
	t6 := `teste
	teste
	teste`

	// %v = valor
	// %T = Tipo
	fmt.Printf("%v --- %T\n", t3, t3)
	fmt.Printf("%v --- %T\n", t4, t4)
	fmt.Printf("%v --- %T\n", t5, t5)
	fmt.Printf("%v --- %T\n", t6, t6)

	fmt.Println(somar(1, 2))
	msgt1("teste")

	// Lib externa
	fmt.Println(guilherme.SomarGuilherme(1, 2))
}

// Tem que colocar o tipo do retorno caso a funcao tenha um retorno
func somar(v1 int, v2 int) int {
	return v1 + v2
}

// Se não retornar nada, nao precisa do tipo de retorno
func msgt1(v1 string) {
	fmt.Println(v1)
}
