package main

import (
	"fmt"
	"log"
	"net/http"
	"hola/internal/handler"
	"hola/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Middleware de logging y tiempo de respuesta
	mux.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(handler.HolaHandler)))

	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
