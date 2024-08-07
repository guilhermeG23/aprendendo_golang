package main

import "net/http"

// Se fizer um curl para o localhost:8888  -> Vai
// retornar o Hello world
func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8888", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow, World!"))
}
