package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// port
const PORT = "8081"

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Printf(
		"Server is running at http://localhost:%s", PORT,
	)

	if err := http.ListenAndServe(
		":"+PORT, r); err != nil {
		fmt.Println(err)
	}
}
