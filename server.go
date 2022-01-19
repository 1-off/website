package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	fmt.Fprintf(w, "Details = %s\n", details.Email)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/hello" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	//this can serve a webpage with js inside to manage the logic and functionalities at the browser level
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// //this manage the webpage logic at the backend level.
	// http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
