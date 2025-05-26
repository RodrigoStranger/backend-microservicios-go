package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Manejador para la ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola! Este es un servidor web básico en Go.")
	})

	// Manejador para la ruta /saludo
	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		nombre := r.URL.Query().Get("nombre")
		if nombre == "" {
			nombre = "Visitante"
		}
		fmt.Fprintf(w, "¡Hola, %s! Bienvenido al servidor Go.", nombre)
	})

	// Configurar el puerto
	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	
	// Iniciar el servidor
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
