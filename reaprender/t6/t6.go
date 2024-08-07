package main

import (
	"encoding/json"
	"fmt"
)

// Manipular a estrutura de retorno json
// No caso, pode requistar pelo nome normal
// porem nos retornos json, ela altera os nome
type V1 struct {
	Name  string `json:"name"`
	Value int
}

// Modo Go de fazer as coisas
// È quase como uma herança
// Vale comentar que ele tmb consegue puxar os métodos
type V2 struct {
	V1
	RandomV bool
}

func (v V1) showName() {
	fmt.Println(v.Name)
}

func main() {
	t1 := V1{
		Name:  "t1",
		Value: 1000,
	}

	fmt.Println(t1.Name)
	fmt.Println(t1.Value)
	t1.showName()

	t2 := V1{"t2", 1000}
	fmt.Println(t2.Name, t2.Value)
	t2.showName()

	// Chamada do incrementa
	t3 := V2{
		V1: V1{
			Name:  "t3",
			Value: 1000,
		},
		RandomV: true,
	}
	fmt.Println(t3.Name, t3.Value, t3.RandomV)
	t3.showName()

	// Converter para json (Slice de byte)
	t3j, err := json.Marshal(t3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t3j)         // Slice de byte
		fmt.Println(string(t3j)) // String do json
	}

	// Popular o t4 com base no json do t3
	// Ele pega e descifra o slice de bytes para struct
	//
	t4 := V2{}
	json.Unmarshal([]byte(t3j), &t4)
	fmt.Println(t4.Name, t4.Value, t4.RandomV)
	t4.showName()
}
