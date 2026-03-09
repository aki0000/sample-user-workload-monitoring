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

		latency := time.Since(start)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = fmt.Fprintf(w, "hello from go on OCP\nlatency: %s\n", latency)
	})

	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
