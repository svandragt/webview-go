package main

import (
	"fmt"
	"net/http"

	"github.com/webview/webview"
)

func main() {
	// Start a goroutine to run the web server
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, Webview!")
		})

		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()

	// Create and run the webview window
	debug := true
	w := webview.New(debug)

	w.Navigate("http://localhost:8080")

	// Run the webview event loop
	w.Run()
}
