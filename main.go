package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(rw, "Form submitted!")
	name := req.FormValue("name")
	tel := req.FormValue("tel")
	fmt.Fprintf(rw, "Name = %s\n", name)
	fmt.Fprintf(rw, "Telephone = %s\n", tel)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(rw, "Welcome!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
