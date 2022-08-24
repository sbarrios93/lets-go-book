package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" dir
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux.Handle() to register the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: The TCP network addess to listen on and the
	// servemux we just created. If http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit.
	// Note that any error return by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
