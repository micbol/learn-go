package main

import (
	"fmt"
	"net/http"
)

// HelloHandler is a Handler
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content := []byte (
	"<h1>Hello, World!</h1>" + 
	"<ul><li><a href='bonjour'>Bonjour</a></li>" + 
	"<li><a href='static/gopher.html'>Gopher</a></li></ul>")
	
	w.Write(content);
	// fmt.Fprintf(w, content)
	
	// fmt.Fprintf(w, "<h1>Hello, World!</h1><ul><li><a href='bonjour'>Bonjour</a></li><li><a href='static/gopher.html'>Gopher</a></li></ul>")
}

// BonjourHandleFunc is a handler function
func BonjourHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bonjour, Monde!</h1>")
}

func main() {

	// Create an instance of the multiplexer
	mux := http.NewServeMux()
	// Create an instance of a handler
	hello := HelloHandler{}
	// Add it to the multiplexer and register it for root "/"
	mux.Handle("/", &hello)
	// Add a handleFunc to the multiplexer and register it for "/bonjour"
	mux.HandleFunc("/bonjour", BonjourHandleFunc)

	// Create a Handler that serves files from folder "static"
	files := http.FileServer(http.Dir("./static"))
	// Add a Handler to the route "static"
	// e.g. http://127.0.0.1:8080/static/gopher.html
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Create a new instance of a server
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting the server and listening...")
	// Start the server
	server.ListenAndServe()

}
