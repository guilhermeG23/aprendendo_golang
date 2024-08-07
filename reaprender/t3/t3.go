package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br")
	//res, err := http.Get("htt1p://google.com.br")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err.Error())
	} else {
		fmt.Println(res.Header)
	}

	//res1, err1 := t3(1, 20)
	res1, err1 := t3(1, 2)

	if err1 != nil {
		log.Fatal(err1.Error())
	} else {
		fmt.Println(res1)
	}

	// Blank return
	// Vc ignora o retorno de algo da funcao
	// Ex: Se voltar 0 é que deu problema, porém sem o err
	// Não sabemos qual é o problema
	res2, _ := t3(1, 20)
	fmt.Println(res2)
}

func t3(v1 int, v2 int) (int, error) {
	res := v1 + v2

	if res >= 10 {
		return 0, errors.New("Maior que 10")
	}

	return res, nil
}
