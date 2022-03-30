package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("carregando homepage ......."))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("carregando area de usuarios......."))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	fmt.Println("Servidor rodando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
