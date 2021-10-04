package main

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -run=TestTicker
func TestTicker(t *testing.T) {
	done := make(chan bool)

	ticker := time.NewTicker(1 * time.Second)

	// run intervally
	x := 0

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			x++
			fmt.Println(x, ". time=> ", t)

		}
	}
}

// go test -v -run=TestTick
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	x := 0
	// run intervally
	for time := range channel {
		x++
		fmt.Println(x, ". time=> ", time)
	}
}
