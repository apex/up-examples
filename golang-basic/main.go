package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := os.Getenv("UP_ADDR")
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from Go")
}
