package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//To mount once the html files to serve
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must((template.ParseFiles(filepath.Join("templates", t.filename))))
	})
	t.templ.Execute(w, nil)
}

func main() {
	//Parse the flag of the port
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("Start server on port: %d", *port)
	//Deliver the html files
	http.Handle("/", &templateHandler{filename: "login.html"})

	//server running
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
