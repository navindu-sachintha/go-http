package main

import (
	"fem/go/http/api"
	"fem/go/http/data"
	"fmt"
	"net/http"
	"text/template"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go http server!"))
	w.WriteHeader(http.StatusOK)
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	html.Execute(w, data.Getall())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error while opening the server")
	}
}
