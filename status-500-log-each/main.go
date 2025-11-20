package main

import (
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

	http.HandleFunc("/", handler)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	log.Printf("[%s] %s %s %s - Returning 500", timestamp, r.Method, r.URL.Path, r.RemoteAddr)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error\n"))
}
