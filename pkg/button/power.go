package button

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/visola/rover/pkg/finalizer"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var powerOff *gpio.ButtonDriver

// InitializePowerOff prepares the power off button
func InitializePowerOff(adaptor *raspi.Adaptor) {
	powerOff = gpio.NewButtonDriver(adaptor, "13")
	powerOff.Start()
	finalizer.AddFinalizer(func() {
		powerOff.Halt()
	})

	powerOff.On(gpio.ButtonRelease, func(_ interface{}) {
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
	})
}
