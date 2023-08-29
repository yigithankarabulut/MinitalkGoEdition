package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var x byte
var i int

func bitstorage(poolsignal os.Signal) {
	if poolsignal == syscall.SIGUSR1 {
		x = x | 1
	}
	i++
	if i == 8 {
		fmt.Printf("%c", x)
		i = 0
		x = 0
	}
	x = x << 1
}

func main() {
	fmt.Printf("Server started with PID: %d\n", os.Getpid())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for {
			poolsignal := <-signalChan
			bitstorage(poolsignal)
		}
	}()

	select {}
}
