package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
)

func init() {
	log.SetHandler(json.Default)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/text/", handleText)
	http.HandleFunc("/json/", handleJSON)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

// text logs.
//
// Note that the indentation here matters, as it
// will be treated as a single log by Up.
func handleText(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request:\n")
	fmt.Printf("  - method: %s\n", r.Method)
	fmt.Printf("  - url: %s\n", r.URL)
	fmt.Printf("  - header:\n")
	for k := range r.Header {
		fmt.Printf("    - %s: %s\n", k, r.Header.Get(k))
	}
	fmt.Fprintln(w, "Hello World from Go")
}

// json logs.
//
// Note that the apex/log json handler already supports
// the format which Up expects in terms of .level, .message
// and .fields properties.
func handleJSON(w http.ResponseWriter, r *http.Request) {
	ctx := log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
		"route":  "handleJSON",
	})

	ctx.Info("doing some work")
	time.Sleep(time.Millisecond * 250)
	ctx.Info("did some work")

	fmt.Fprintln(w, "Hello World from Go")
}
