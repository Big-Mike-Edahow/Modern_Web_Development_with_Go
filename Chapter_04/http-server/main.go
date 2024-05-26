// main.go
// Simple HTTP Server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<h1>The Go Programming Language</h1><h2>Hello World! 🌎</h2><h3>Welcome to Big Mike's Website! 🌐</h3>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
