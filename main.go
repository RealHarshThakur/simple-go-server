package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed static/giphy.webp
var indexPage []byte

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server on port", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(indexPage)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
