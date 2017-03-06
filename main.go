package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logger(server http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s - %s", r.Method, r.URL, r.RemoteAddr)
		server.ServeHTTP(w, r)
	})
}

var hostname, _ = os.Hostname()

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! (%s)\n", hostname)
	})
	log.Println("starting server on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", logger(server)))
}
