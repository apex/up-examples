package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tj/go/http/response"
)

// Pet model.
type Pet struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

// faux db.
var db = []Pet{
	{
		Name:    "Tobi",
		Species: "Ferret",
		Age:     3,
	},
	{
		Name:    "Loki",
		Species: "Ferret",
		Age:     2,
	},
	{
		Name:    "Jane",
		Species: "Ferret",
		Age:     5,
	},
	{
		Name:    "Luna",
		Species: "Cat",
		Age:     5,
	},
	{
		Name:    "Manny",
		Species: "Cat",
		Age:     5,
	},
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/pets/", pets)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func pets(w http.ResponseWriter, r *http.Request) {
	response.OK(w, db)
}
