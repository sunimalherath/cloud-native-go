package main

import (
	"fmt"
	"github.com/sunimalherath/cloud-native-go/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandlFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "welcome to the cloud native go...")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("content-type", "text/plain")
	fmt.Fprintf(w, message)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	}
	return ":" + port
}
