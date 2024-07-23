package main

import (
	"fmt"
	"net/http"
	"ascii-art-web/ascii-art"
)

func main() {
	http.HandleFunc("/", asciiart.ServeIndexPage)
	http.HandleFunc("/ascii-art", asciiart.FormHandler)

	fmt.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
