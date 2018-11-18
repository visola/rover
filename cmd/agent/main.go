package main

import (
	"github.com/visola/rover/pkg/adaptor"
	"github.com/visola/rover/pkg/api"
	"github.com/visola/rover/pkg/finalizer"
)

func main() {
	adaptor.Start()
	finalizer.RegisterHook()
	api.Start()
}
