package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
)

var views = template.Must(template.ParseGlob("views/*.html"))

func main() {
	log.SetHandler(logfmt.New(os.Stdout))
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	name := cookie(r, "name")
	email := cookie(r, "email")

	views.ExecuteTemplate(w, "index.html", struct {
		Name  string
		Email string
	}{
		Name:  name,
		Email: email,
	})

	w.Header().Set("Content-Type", "text/html")
}

func submit(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	log.WithFields(log.Fields{
		"name":  name,
		"email": email,
	}).Info("submit")

	http.SetCookie(w, &http.Cookie{
		Name:  "name",
		Value: name,
	})

	redirectBack(w, r)
}

func redirectBack(w http.ResponseWriter, r *http.Request) {
	url := r.Header.Get("Referer")
	http.Redirect(w, r, url, http.StatusFound)
}

func cookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	return c.Value
}
