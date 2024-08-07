package main

import "fmt"

type Estrutura struct {
	Name string
}

func (e Estrutura) showname() {
	fmt.Println(e.Name)
}

func main() {
	t1, t2 := ft1(1, 2)
	fmt.Println(t1, t2)

	t3 := ft2(1, 2)
	fmt.Println(t3)

	t4 := ft3(1, 2, 4, 5, 6, 7, 8)
	fmt.Println(t4)

	// funcao anonima
	result := func() int {
		return 10
	}

	// Se chamar sem o (), ele chama a
	// referencia de memoria da funcao,
	// Se usar o (), ele chama o valor de
	// retorno da funcao
	fmt.Println(result)
	fmt.Println(result())

	// Chamada de estrutura
	t5 := Estrutura{
		Name: "t5",
	}

	t5.showname()
}

func ft1(v1 int, v2 int) (int, string) {
	return v1 + v2, "ft1"
}

// Retorno nomeado
func ft2(v1 int, v2 int) (result int) {
	result = v1 + v2
	return
}

// Funcoes variadicas
// Aceita N entradas na funcao
func ft3(x ...int) int {
	result := 0

	// Key, value
	for _, v := range x {
		result += v
	}

	return result
}
