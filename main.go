package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		latency := time.Since(start)
		_, _ = fmt.Fprintf(w, "hello from go on OCP\nlatency: %s\n", latency)
	})

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           nil, // default mux
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("server listening on :%s", port)
	log.Fatal(server.ListenAndServe())
}
