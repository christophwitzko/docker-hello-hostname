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

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	bindAddr := ":5000"
	if port := os.Getenv("PORT"); port != "" {
		bindAddr = ":" + port
	}
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! (%s)\n", hostname)
	})
	log.Printf("starting server on port %s...", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, logger(server)))
}
