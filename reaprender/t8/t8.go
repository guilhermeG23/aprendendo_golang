package main

import "fmt"

func somaInt(m map[string]int64) int64 {
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Gerar um generic para multiplos tipos de entradas
func somaGeneric[T int64 | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Gerar um generic para multiplos tipos de entradas e com
// um tipo customizado só para ela
type mycustomtype interface {
	~int64 | ~float64 | ~float32 // Esse ~ é uma aproximação dos tipos
}

func somaGeneric1[T mycustomtype](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Outro tipo
type pastel int64

// Permitir comparar valroes aleatorios de entrada
// Más só permite comparar
func somaAlea[T comparable](n1 T, n2 T) T {
	if n1 == n2 {
		return n1
	}

	return n2
}

// Teste de generics
func main() {
	// Usually
	fmt.Println(somaInt(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(somaFloat(map[string]float64{"a": 1.1, "b": 2, "c": 3}))

	// Generic
	fmt.Println(somaGeneric(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(somaGeneric(map[string]float64{"a": 1.1, "b": 2, "c": 3}))

	// Generic with interface
	fmt.Println(somaGeneric1(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(somaGeneric1(map[string]float64{"a": 1.1, "b": 2, "c": 3}))

	var t1, t2, t3 pastel
	t1 = 1
	t2 = 1
	t3 = 1
	// Ele da ruim se colocar direto pq os tipos conflitam
	//fmt.Println(somaGeneric1(map[string]pastel{"a": t1, "b": t2, "c": t3}))
	fmt.Println(somaGeneric1(map[string]pastel{"a": t1, "b": t2, "c": t3}))

	// Comparar entradas diferente
	fmt.Println(somaAlea(1, 1))
	fmt.Println(somaAlea(1, 1.0))
	fmt.Println(somaAlea(1, 2.0))
}
