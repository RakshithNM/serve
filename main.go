package main

import (
	"fmt"
	"log"
	"net/http"
)

const tpl = `
    RequestURI: %v
    Host:       %v
    Form:       %v
        some    %v
`

func reflectForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, tpl, r.RequestURI, r.Host, r.Form, r.Form.Get("some"))
}

func main() {
	http.HandleFunc("/", reflectForm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
