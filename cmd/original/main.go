package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Uses package-level log (no shared structured logger)
		log.Printf("started method=%s path=%s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("completed method=%s path=%s duration=%s", r.Method, r.URL.Path, time.Since(start))
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("home handler hit")
	w.Write([]byte("Welcome (original, no DI)\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	handler := loggingMiddleware(mux)

	log.Println("starting server on :4000 (original)")
	err := http.ListenAndServe(":4000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
