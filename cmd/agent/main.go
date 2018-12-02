package main

import (
	"github.com/visola/rover/pkg/adaptor"
	"github.com/visola/rover/pkg/api"
	"github.com/visola/rover/pkg/button"
	"github.com/visola/rover/pkg/finalizer"
	"github.com/visola/rover/pkg/light"
)

func main() {
	adaptor.Start()
	light.InitializeFrontLight(adaptor.RPi)
	button.InitializePowerOff(adaptor.RPi)
	finalizer.RegisterHook()
	api.Start()
}
