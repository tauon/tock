package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	secondsBetweenTicks = 10
)

func main() {
	quitSignal := make(chan os.Signal, 1)
	quit := make(chan bool, 1)
	ticker := time.NewTicker(secondsBetweenTicks * time.Second)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quitSignal
		// cleanup
		quit <- true
	}()

	go func() {
		for _ = range ticker.C {
			// do stuff
		}
	}()
	<-quit // wait for quit
	fmt.Println("bye!")
}
