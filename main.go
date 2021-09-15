package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello From Snippet Box"))
}

func getSpecificSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From specific Snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Hello From Create Snippet "))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", getSpecificSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Serve App on Port 4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
