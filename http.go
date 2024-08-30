package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, you are on Home!")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, you are on Info!")
}

func main() {
	log.Println("Starting a http server...")

	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("Started server on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Error running the http server!")
	}
}
