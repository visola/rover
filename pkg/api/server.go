package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/visola/rover/pkg/wheel"
)

// Start starts the REST API server and register all endpoints
func Start() {
	r := mux.NewRouter()
	wheel.RegisterEndpoints(r)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", r)
}
