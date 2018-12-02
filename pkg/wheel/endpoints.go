package wheel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/visola/rover/pkg/adaptor"
	"github.com/visola/rover/pkg/light"
)

var wheelDriver *Driver

// Movement VO to store car movement
type Movement struct {
	XAxis int
	YAxis int
}

// RegisterEndpoints registers all endpoints associated with the wheel
func RegisterEndpoints(router *mux.Router) {
	fmt.Println("Registering wheel endpoints...")
	wheelDriver = NewDriver(adaptor.RPi)
	router.HandleFunc("/wheels", setWheelsMovement).Methods(http.MethodPut)
}

func setWheelsMovement(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var movement Movement

	err := decoder.Decode(&movement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Sorry, something went wrong: '%s'\n", err)))
	}

	if movement.YAxis < 0 {
		light.FrontLightOn()
	} else {
		light.FrontLightOff()
	}

	wheelDriver.Move(movement.YAxis)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
