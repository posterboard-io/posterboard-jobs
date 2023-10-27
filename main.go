package main



import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/message", messageHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// This will serve our HTML page (we'll create this soon)
	http.ServeFile(w, r, "index.html")
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// This will return a simple message when requested
	fmt.Fprint(w, "Hello from Go!")
}
