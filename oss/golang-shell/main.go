package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/tj/go/http/response"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", command)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func command(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response.BadRequest(w)
		return
	}

	cmd := exec.Command("sh", "-c", r.FormValue("command"))
	b, err := cmd.CombinedOutput()
	if err != nil {
		response.InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(b)
}
