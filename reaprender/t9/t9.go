package main

import (
	"context"
	"fmt"
	"time"
)

// Trabalhando com contexto
// Trabalha em arvore, o "pai" afeta os "filhos"
func main() {
	// Contexto pai
	ctx := context.Background()
	// Retorna a funcao de cancelamento (Todo o filho de determinado pai morre)
	ctx, cancel := context.WithCancel(ctx)

	// Cancel por timeout
	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	// Espera tudo ser executado para dar um cancel
	defer cancel()

	// Simula um cancelamento
	// Ou um processamento paralelo
	go func() {
		time.Sleep(time.Second * 5) // Se for 5 ou mais, significa que a funcao tem mais chance de rodar e terminar certo
		//time.Sleep(time.Second * 3) // Aqui vai ser certeza que vai cancelar até de processar a funcao de baixo
		cancel()
	}()

	// processa a reserva
	bookHotel(ctx)
}

// Por boas prática o contexto tem sempre que ser
// o primeiro parametro nas funcoes que o usarem
func bookHotel(ctx context.Context) {
	// Espera alguma condição bater para retornar a acao
	select {
	case <-ctx.Done():
		fmt.Println("Tem excedido")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado")
	}
}
