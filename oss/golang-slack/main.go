package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apex/httpstat"
	"github.com/dustin/go-humanize"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// parse form body
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// grab values from form
	team := r.Form.Get("team_domain")
	user := r.Form.Get("user_name")
	cmd := r.Form.Get("command")
	args := strings.Fields(r.Form.Get("text"))

	if len(args) == 0 {
		fmt.Fprintln(w, "url is required")
		return
	}

	// log stuff
	fmt.Printf("%s %v from user %s (%s)\n", cmd, args, user, team)

	// perform the request
	res, err := httpstat.Request("GET", args[0], nil, nil)
	if err != nil {
		log.Printf("url required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// write timing information
	fmt.Fprintf(w, "Status %d\n", res.Status())
	fmt.Fprintf(w, "Size %s\n", humanize.Bytes(uint64(res.BodySize())))
	fmt.Fprintf(w, "Timing:")
	fmt.Fprintf(w, "  - DNS %s\n", res.TimeDNS())
	fmt.Fprintf(w, "  - Connect %s\n", res.TimeConnect())
	fmt.Fprintf(w, "  - TLS %s\n", res.TimeTLS())
	fmt.Fprintf(w, "  - Wait %s\n", res.TimeWait())
	fmt.Fprintf(w, "  - Response %s\n", res.TimeResponse(time.Now()))
	fmt.Fprintf(w, "  - Download %s\n", res.TimeDownload(time.Now()))
	fmt.Fprintf(w, "  - Total %s\n", res.TimeTotalWithRedirects(time.Now()))
}
