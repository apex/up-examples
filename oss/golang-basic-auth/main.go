package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.Handle("/", protect(http.HandlerFunc(hello)))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		match := user == "tobi" && pass == "ferret"

		if !ok || !match {
			w.Header().Set("WWW-Authenticate", `Basic realm="Ferret Land"`)
			http.Error(w, `Unauthorized: Use "tobi" and "ferret" :)`, http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Access to the secret ferret land granted!")
}
