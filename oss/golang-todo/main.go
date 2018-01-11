package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/pat"
	"github.com/hoisie/mustache"
	"github.com/tj/go/http/response"
)

// items is the list of items, note that you should
// use a database such as DynamoDB, as you will not
// always hit the same Lambda container (items may disappear),
// but for the purpose of this illustration it works.
var items []string

func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/", index)
	app.Post("/submit", submit)
	log.Fatal(http.ListenAndServe(addr, app))
}

// index page.
func index(w http.ResponseWriter, r *http.Request) {
	render(w, "index", struct {
		Items []string
	}{
		Items: items,
	})
}

// submit handler.
func submit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response.BadRequest(w)
		return
	}

	item := r.FormValue("item")
	items = append(items, item)

	back(w, r)
}

// render a view.
func render(w http.ResponseWriter, name string, data interface{}) {
	path := filepath.Join("views", name+".html")

	t, err := mustache.ParseFile(path)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, t.Render(data))
}

// back redirects the user back.
func back(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
