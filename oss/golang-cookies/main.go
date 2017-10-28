package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
)

var views = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	addr := ":" + os.Getenv("PORT")
	s := pat.New()
	s.Get("/", index)
	s.Post("/", submit)
	log.Fatal(http.ListenAndServe(addr, s))
}

func index(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Cookie("name")

	w.Header().Set("Content-Type", "text/html")

	views.ExecuteTemplate(w, "index.html", struct {
		Name *http.Cookie
	}{
		Name: name,
	})
}

func submit(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	http.SetCookie(w, &http.Cookie{
		Name:  "name",
		Value: name,
	})

	back := r.Header.Get("Referer")
	http.Redirect(w, r, back, http.StatusFound)
}
