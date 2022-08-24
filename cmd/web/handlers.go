package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the path to the two files. It's important
	// that the file containing the base template must be the first file in the slice
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there''s an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to
	// Execute() represents any dynamic data that we want to pass in, which for now
	// we'll leave as nil
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query
	// string and try to convert it to an integer. Return error
	// if it can be converted to int or if the val < 1

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check if the request is POST
	if r.Method != http.MethodPost {
		// Broadcast allowed methods
		w.Header().Set("Allow", http.MethodPost)
		// If its not, send a 405 status and use w.Wrote() to write
		// a "method not allowed" response body.
		// ===== like this
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// ===== or
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
