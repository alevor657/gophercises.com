package handlers

import (
	"net/http"

	"github.com/alevor657/gophercises/handlers/utils"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		redirect, exists := pathsToUrls[requestPath]

		if exists {
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yaml := utils.ParseYaml(yml)
	return nil, nil
}
