package light

import (
	"github.com/visola/rover/pkg/finalizer"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var frontLightDriver *gpio.LedDriver

// InitializeFrontLight prepares the front light to be used
func InitializeFrontLight(adaptor *raspi.Adaptor) {
	frontLightDriver = gpio.NewLedDriver(adaptor, "11")
	finalizer.AddFinalizer(func() {
		frontLightDriver.Halt()
	})
}

// FrontLightOff turn off the front light
func FrontLightOff() {
	frontLightDriver.Off()
}

// FrontLightOn turn on the front light
func FrontLightOn() {
	frontLightDriver.On()
}
