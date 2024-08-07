package main

import "fmt"

func main() {
	t1 := 10
	fmt.Println(&t1) // Printa o endereço de memoria

	var t2 *int = &t1
	fmt.Println(t2)  // Printa o endereco de memoria
	fmt.Println(*t2) // Printa o valor que tem o endereco de memoria

	// Alterando o valor de t1 via ponteiro
	*t2 = 20
	fmt.Println(&t1) // Printa o endereço de memoria
	fmt.Println(t1)  // Printa o valor de t1
	fmt.Println(t2)  // Printa o endereco de memoria do t2 (t1)
	fmt.Println(*t2) // Printa o valor da referencia

	changevalue(&t1)
	fmt.Println(t1) // Printa o valor de t1

	// Escopo
	t3 := Teste1{
		Name: "Test1",
	}
	t3.showName()
	fmt.Println(t3.Name)
	t3.showName1()
	fmt.Println(t3.Name)
}

type Teste1 struct {
	Name string
}

// Funcoes de struct são métodos

// Altera o valor, porém limitado a esse escopo
func (t Teste1) showName() {
	t.Name = "pastel"
	fmt.Println(t.Name)
}

// Altera o valor com base na referencia de memoria
func (t *Teste1) showName1() {
	t.Name = "pastel"
	fmt.Println(t.Name)
}
func changevalue(v1 *int) {
	*v1 = 100
}
