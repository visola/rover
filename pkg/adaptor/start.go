package adaptor

import (
	"fmt"

	"github.com/visola/rover/pkg/finalizer"
	"gobot.io/x/gobot/platforms/raspi"
)

// RPi is the RaspberryPi Adaptor to use
var RPi *raspi.Adaptor

// Start the RPi adaptor
func Start() {
	fmt.Println("Starting application...")
	RPi = raspi.NewAdaptor()
	finalizer.AddFinalizer(func() {
		RPi.Finalize()
	})
}
