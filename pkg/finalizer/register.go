package finalizer

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var finalizers = make([]func(), 0)

// AddFinalizer register a callback to be called
func AddFinalizer(toCall func()) {
	finalizers = append(finalizers, toCall)
}

// RegisterHook registers the signal hook to call the finailzers
func RegisterHook() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			fmt.Println("Closing RPi adaptor...")
			for i := len(finalizers) - 1; i >= 0; i-- {
				fmt.Printf("Finalizing %d\n", i)
				finalizers[i]()
			}
			os.Exit(0)
		}
	}()
}
