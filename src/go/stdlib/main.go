package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	webport := "8080"
	http.HandleFunc("/", homepage)
	err := http.ListenAndServe(":"+webport, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
