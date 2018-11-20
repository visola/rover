package wheel

import (
	"github.com/visola/rover/pkg/finalizer"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	frequency  = 63
	resolution = 4096 // 12 bits

	// 1 second in microseconds
	timePerCycle = 1000000 / frequency / resolution

	// Stop rotating at the 1.5s
	stop = 1500 / timePerCycle

	// Max speed CW -> 1s
	cw = 1000 / timePerCycle

	// Max speed CCW -> 2s
	ccw = 2000 / timePerCycle
)

// Driver drives the wheels
// Assumes the motors are set in the following position:
// - Front -
// 2       1
//
// 0       3
// - Back  -
type Driver struct {
	driver *i2c.PCA9685Driver
}

// NewDriver creates a new wheel driver from the RaspberryPi adapter
func NewDriver(adaptor *raspi.Adaptor) *Driver {
	wheelDriver := &Driver{
		driver: i2c.NewPCA9685Driver(adaptor),
	}
	wheelDriver.driver.Start()
	wheelDriver.driver.SetPWMFreq(50)
	finalizer.AddFinalizer(func() {
		wheelDriver.driver.Halt()
	})
	return wheelDriver
}

// MoveBackwards drives backwards
func (w *Driver) MoveBackwards() {
	w.setLeftSide(0, ccw)
	w.setRightSide(0, cw)
}

// MoveForward drives forward
func (w *Driver) MoveForward() {
	w.setLeftSide(0, cw)
	w.setRightSide(0, ccw)
}

// Stop stops all wheels
func (w *Driver) Stop() {
	w.setLeftSide(0, stop)
	w.setRightSide(0, stop)
}

func (w *Driver) setLeftSide(start uint16, stop uint16) {
	w.driver.SetPWM(0, start, stop)
	w.driver.SetPWM(2, start, stop)
}

func (w *Driver) setRightSide(start uint16, stop uint16) {
	w.driver.SetPWM(1, start, stop)
	w.driver.SetPWM(3, start, stop)
}
