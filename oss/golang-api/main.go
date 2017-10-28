package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/gorilla/pat"
)

// pets database-ish
var pets = make(map[string]struct{})

// use JSON logging when run by Up (including `up start`).
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(json.Default)
	}
}

// setup.
func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/", get)
	app.Post("/", post)
	app.Delete("/{name}", del)
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

// curl http://localhost:3000/
func get(w http.ResponseWriter, r *http.Request) {
	log.Info("list pets")

	if len(pets) == 0 {
		fmt.Fprintf(w, "no pets")
		return
	}

	for name := range pets {
		fmt.Fprintf(w, "- %s\n", name)
	}
}

// curl -d Tobi http://localhost:3000/
// curl -d Loki http://localhost:3000/
// curl -d Jane http://localhost:3000/
func post(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	name := string(b)
	pets[name] = struct{}{}
	log.WithField("name", name).Info("add pet")
	fmt.Fprintf(w, "welcome to the family %s!\n", name)
}

// curl -X DELETE http://localhost:3000/Tobi
// curl -X DELETE http://localhost:3000/Loki
// curl -X DELETE http://localhost:3000/Jane
func del(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":name")
	log.WithField("name", name).Info("remove pet")
	delete(pets, name)
	fmt.Fprintf(w, "removed %s!\n", name)
}
