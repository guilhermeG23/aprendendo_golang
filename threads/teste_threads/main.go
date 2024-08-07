package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := range x {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { // Thread 1
	go contador(10) // Thread 2
	go contador(10) // Thread 3
	contador(10)    // Por esse motivo esse cara nao tem go
}
