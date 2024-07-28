package main

import (
	"net/http"
	"os"

	"github.com/sousaprogramador/api-trasnfer-go/internal/money"
)

func main() {
	http.HandleFunc("/transfers", money.TransferHandler)
	http.HandleFunc("/users/", money.UsersHandler)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, nil)
}
