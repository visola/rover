package button

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/visola/rover/pkg/finalizer"
	"github.com/visola/rover/pkg/light"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var powerOff *gpio.ButtonDriver

var powerOffDown = false
var powerOffPressedAt time.Time

// InitializePowerOff prepares the power off button
func InitializePowerOff(adaptor *raspi.Adaptor) {
	powerOff = gpio.NewButtonDriver(adaptor, "13")
	finalizer.AddFinalizer(func() {
		powerOff.Halt()
	})

	go func() {
		for {
			if powerOffDown {
				diff := time.Now().Sub(powerOffPressedAt).Seconds()
				if diff > 5 {
					light.FrontLightOff()
					shutdown()
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	powerOff.On(gpio.ButtonPush, func(_ interface{}) {
		powerOffDown = false
		light.FrontLightOff()
	})

	powerOff.On(gpio.ButtonRelease, func(_ interface{}) {
		powerOffPressedAt = time.Now()
		powerOffDown = true
		light.FrontLightOn()
	})

	powerOff.Start()
}

func shutdown() {
	fmt.Println("Powering off...")

	cmd := exec.Command("sudo", "shutdown", "-h", "now")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Standard out:\n%q\n", stdOut.String())
	fmt.Printf("Standard err:\n%q\n", stdErr.String())

	os.Exit(0)
}
