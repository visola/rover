package wheel

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/visola/rover/pkg/adaptor"
)

var wheelDriver *Driver

// RegisterEndpoints registers all endpoints associated with the wheel
func RegisterEndpoints(router *mux.Router) {
	fmt.Println("Registering wheel endpoints...")
	wheelDriver = NewDriver(adaptor.RPi)
	router.HandleFunc("/wheels/backwards", driverBackwards)
	router.HandleFunc("/wheels/forward", driverForward)
	router.HandleFunc("/wheels/stop", stopDriving)
}

func driverBackwards(w http.ResponseWriter, req *http.Request) {
	wheelDriver.MoveBackwards()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func driverForward(w http.ResponseWriter, req *http.Request) {
	wheelDriver.MoveForward()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func stopDriving(w http.ResponseWriter, req *http.Request) {
	wheelDriver.Stop()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
