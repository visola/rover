package api

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/visola/rover/pkg/wheel"
)

// Start starts the REST API server and register all endpoints
func Start() {
	router := mux.NewRouter()
	wheel.RegisterEndpoints(router)

	box := packr.NewBox("../../web")
	router.Handle("/{file:.*}", http.FileServer(box))

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", router)
}
