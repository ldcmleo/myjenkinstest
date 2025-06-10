package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola, mundo!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Servidor escuchando en http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
