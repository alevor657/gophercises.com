package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alevor657/gophercises/handlers"
)

func main() {
	mux := newMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/t1": "https://google.se",
		"/t2": "https://yandex.ru",
	}

	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	fmt.Println("Starting the server on :8080")
	// http.ListenAndServe(":8080", yamlHandler)

	log.Fatal(http.ListenAndServe(":8080", mapHandler))
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
