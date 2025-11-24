package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	serviceurl := os.Getenv("CHOREO_TESTING_SERVICEURL")
	choreoapikey := os.Getenv("CHOREO_TESTING_CHOREOAPIKEY")

	if serviceurl == "" {
		fmt.Println("NO URL SET")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(serviceurl)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer response.Body.Close()

		w.WriteHeader(response.StatusCode)
		r.Header.Add("Choreo-API-Key", choreoapikey)

		body, err := io.ReadAll(response.Body)

		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			return
		}

		w.Write(body)
	})

	log.Printf("Starting proxy server on :8080, forwarding to %s", serviceurl)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
