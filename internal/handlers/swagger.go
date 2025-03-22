package handlers

import (
	"net/http"
)

func SwaggerHandler(swaggerPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	}
}