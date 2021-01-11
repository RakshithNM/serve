package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const tpl = `
    RequestURI: %v
    Host:       %v
    Form:       %v
        some    %v
`

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, tpl, r.RequestURI, r.Host, r.Form, r.Form.Get("some"))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, filepath.Join(wd, r.URL.Path))
}

func main() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/form", serveForm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
