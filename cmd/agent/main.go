package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		data := fmt.Sprintf("Time is now: %s\n", time.Now().Format("20060102150405"))
		ioutil.WriteFile("/home/pi/time.txt", []byte(data), 0644)
		time.Sleep(1 * time.Second)
	}
}
