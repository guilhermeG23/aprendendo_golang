package main

import (
	"fmt"
	"time"
)

func contador(qtd int) {
	for i := range qtd {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { // Thr1
	go contador(10) // Thr2
	go contador(10) // Thr3
	contador(10)
}
