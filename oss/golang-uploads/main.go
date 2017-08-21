package main

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/dustin/go-humanize"
)

var funcs = template.FuncMap{
	"humanize_bytes": humanize.Bytes,
}

var views = template.Must(template.New("").Funcs(funcs).ParseGlob("views/*.html"))

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
	w.Header().Set("Content-Type", "text/html")
	views.ExecuteTemplate(w, "index.html", nil)
}

func submit(w http.ResponseWriter, r *http.Request) {
	file, hdr, err := r.FormFile("image")
	if err != nil {
		log.WithError(err).Error("parsing form")
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithError(err).Error("reading file")
		http.Error(w, "Error reading file.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	views.ExecuteTemplate(w, "index.html", struct {
		Name  string
		Size  uint64
		Type  string
		Image string
	}{
		Name:  hdr.Filename,
		Size:  uint64(hdr.Size),
		Type:  hdr.Header.Get("Content-Type"),
		Image: base64.StdEncoding.EncodeToString(b),
	})
}
