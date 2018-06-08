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
	first, _ := r.Cookie("first")
	last, _ := r.Cookie("last")

	w.Header().Set("Content-Type", "text/html")

	views.ExecuteTemplate(w, "index.html", struct {
		First string
		Last  string
	}{
		First: first.String(),
		Last:  last.String(),
	})
}

func submit(w http.ResponseWriter, r *http.Request) {
	first := r.FormValue("first")
	last := r.FormValue("last")

	http.SetCookie(w, &http.Cookie{
		Name:  "first",
		Value: first,
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "last",
		Value: last,
	})

	back := r.Header.Get("Referer")
	http.Redirect(w, r, back, http.StatusFound)
}
