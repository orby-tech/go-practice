package main

import (
	"fmt"
	"net/http"
)

// port
const PORT = "8081"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf(
		"Server is running at http://localhost:%s", PORT,
	)

	if err := http.ListenAndServe(
		":"+PORT, nil); err != nil {
		fmt.Println(err)
	}

}
