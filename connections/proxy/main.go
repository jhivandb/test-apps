package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("REQUEST_URL")

	if url == "" {
		log.Fatal("NO URL SET")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(url)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer response.Body.Close()

		w.WriteHeader(response.StatusCode)

		body, err := io.ReadAll(response.Body)

		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			return
		}

		w.Write(body)
	})

	log.Printf("Starting proxy server on :8080, forwarding to %s", url)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
