package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, string(dump))
}
