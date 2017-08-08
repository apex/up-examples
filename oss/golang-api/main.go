package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
)

// pets database-ish
var pets = make(map[string]struct{})

func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/", get)
	app.Post("/", post)
	app.Delete("/{name}", del)
	log.Fatal(http.ListenAndServe(addr, app))
}

// curl :3000/
func get(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")

	if len(pets) == 0 {
		fmt.Fprintf(w, "no pets")
		return
	}

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
