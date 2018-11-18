package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	start = 25
	stop  = 310
	cw    = stop - start
	ccw   = stop + start
)

var buttonState = false

func main() {
	fmt.Println("Starting application...")
	r := raspi.NewAdaptor()
	motorDriver := i2c.NewPCA9685Driver(r)

	work := func() {
		motorDriver.Start()
		motorDriver.SetPWMFreq(50)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("W: Move forward\nS: Stop\nX: Backwards")

		for {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if text == "w" {
				moveForward(motorDriver)
			}

			if text == "s" {
				moveBackward(motorDriver)
			}

			if text == "x" {
				stopMotors(motorDriver)
			}
		}
	}

	robot := gobot.NewRobot("rover",
		[]gobot.Connection{r},
		[]gobot.Device{motorDriver},
		work,
	)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		robot.Stop()
		os.Exit(0)
	}()

	robot.Start()
}

// 2     1
//
// 0     3

func moveBackward(motorDriver *i2c.PCA9685Driver) {
	motorDriver.SetPWM(0, 0, ccw)
	motorDriver.SetPWM(2, 0, ccw)

	motorDriver.SetPWM(1, 0, cw)
	motorDriver.SetPWM(3, 0, cw)
}

func moveForward(motorDriver *i2c.PCA9685Driver) {
	motorDriver.SetPWM(0, 0, cw)
	motorDriver.SetPWM(2, 0, cw)

	motorDriver.SetPWM(1, 0, ccw)
	motorDriver.SetPWM(3, 0, ccw)
}

func stopMotors(motorDriver *i2c.PCA9685Driver) {
	motorDriver.SetPWM(0, 0, stop)
	motorDriver.SetPWM(1, 0, stop)
	motorDriver.SetPWM(2, 0, stop)
	motorDriver.SetPWM(3, 0, stop)
}
