package main

import (
	"os"
	"strconv"
	"syscall"
	"time"
)

func signalcheck(pid int, data string) {
	j := 0
	for j < len(data) {
		i := 7
		bit := data[j]
		for i >= 0 {
			if (bit >> uint(i) & 1) == 1 {
				_ = syscall.Kill(pid, syscall.SIGUSR1)
			} else {
				_ = syscall.Kill(pid, syscall.SIGUSR2)
			}
			time.Sleep(200 * time.Microsecond)
			i--
		}
		j++
	}
}

func main() {
	args := os.Args
	if len(args) == 3 {
		pid, _ := strconv.Atoi(args[1])
		data := args[2]
		signalcheck(pid, data)
	}
}
