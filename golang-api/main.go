package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/segmentio/go-env"
)

// pets database-ish
var pets = make(map[string]struct{})

func main() {
	app := pat.New()
	app.Get("/", get)
	app.Post("/", post)
	app.Delete("/{name}", del)
	log.Fatal(http.ListenAndServe(env.MustGet("UP_ADDR"), app))
}

// curl :3000/
func get(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	for name := range pets {
		fmt.Fprintf(w, "- %s\n", name)
	}
}

// curl -d Tobi :3000/
// curl -d Loki :3000/
// curl -d Jane :3000/
func post(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	name := string(b)
	pets[name] = struct{}{}
	log.Printf("POST / %q", name)
	fmt.Fprintf(w, "welcome to the family %q!\n", name)
}

// curl -X DELETE :3000/Tobi
// curl -X DELETE :3000/Loki
// curl -X DELETE :3000/Jane
func del(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":name")
	delete(pets, name)
}
