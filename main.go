package main

import (
	"fmt"
	"net/http"

	asciiart "ascii-art-web/handler"
)

func main() {
	http.HandleFunc("/", asciiart.GetHandler)
	http.HandleFunc("/ascii-art",asciiart.FormHandler)

	fmt.Println("Starting server on :http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
