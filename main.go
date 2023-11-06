package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/skratchdot/open-golang/open"
)

func CastReceiverHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	fmt.Printf("url: %s\n", url)
	if url == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("Navigating to %s\n", url)

	// Check if the URL is a YouTube URL
	if strings.Contains(url, "youtube.com") {
		fmt.Printf("YouTube URL detected\n")
		// Modify the URL to use the a full screen YouTube player
		url = strings.Replace(url, "watch?v=", "embed/", 1)

		// Add the autoplay parameter
		url = url + "?autoplay=1"
	}

	// Open the default web browser and navigate to the modified URL
	err := open.Run(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open the browser: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Navigating to %s", url)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cast", CastReceiverHandler).Methods("GET")

	serverAddr := ":8080"
	fmt.Printf("Listening on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
