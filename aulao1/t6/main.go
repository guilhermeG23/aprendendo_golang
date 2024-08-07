package main

import "fmt"

// Interface
type Veiculo interface {
	andarCarro()
}

// Implements é implicito
type Carro struct {
	Modelo string
	Ano    int
}

func (c Carro) andarCarro() {
	fmt.Println(c.Modelo, c.Ano)
}

func chamarOAndar(c Carro) {
	c.andarCarro()
}

func main() {
	c1 := Carro{Modelo: "t1", Ano: 1}
	chamarOAndar(c1)
}
