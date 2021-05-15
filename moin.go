package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler was called")
	fmt.Fprintf(w, "Moin moin!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
