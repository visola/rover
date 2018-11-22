package wheel

import (
	"fmt"
	"math"

	"github.com/visola/rover/pkg/finalizer"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	frequency  = 50.
	resolution = 4096. // 12 bits

	// 1 second in microseconds
	timePerCycle = 1000000. / frequency / resolution
)

var (
	// Stop rotating at the 1.5s
	stop = uint16(math.Round(1500. / timePerCycle))

	// Max speed CW -> 1s
	cw = uint16(math.Round(1000. / timePerCycle))

	// Max speed CCW -> 2s
	ccw = uint16(math.Round(2000. / timePerCycle))

	// Difference from max speed forward and stop
	diff = ccw - stop
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
	fmt.Printf("Stop: %d, CCW: %d, CW: %d\n", stop, ccw, cw)
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

// Move sets the speed between 100 and -100
func (w *Driver) Move(speed int) {
	speedInCycles := uint16(speed * int(diff) / 100)
	w.setRightSide(0, stop+speedInCycles)
	w.setLeftSide(0, stop-speedInCycles)
}

func (w *Driver) setLeftSide(start uint16, stop uint16) {
	w.driver.SetPWM(0, start, stop)
	w.driver.SetPWM(2, start, stop)
}

func (w *Driver) setRightSide(start uint16, stop uint16) {
	w.driver.SetPWM(1, start, stop)
	w.driver.SetPWM(3, start, stop)
}
