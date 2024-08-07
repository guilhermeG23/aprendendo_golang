package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou request")
	defer log.Println("Fim do request")

	select {
	case <-time.After(time.Second * 5):
		fmt.Println(w, "Requisicao processada")
		w.Write([]byte("Requisicao processada"))
	case <-ctx.Done():
		log.Println("Requisicao cancelada")
		http.Error(w, "Requisicao cancelada", http.StatusRequestTimeout)
	}
}
