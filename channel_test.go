package main

import (
	"fmt"
	"testing"
	"time"
)

//  go test -v -run=TestCreateChannel
func TestCreateChannel(T *testing.T) {
	chn := make(chan string)
	defer close(chn)

	go func() { //anonymouse function
		time.Sleep(2 * time.Second)
		// channel must be taken so the code doesnt block
		chn <- "hello channel"
		fmt.Println("[done donk!]")
	}() //auto run

	// channel must be defined & sent so the code doesnt asleep - deadlock
	data := <-chn
	fmt.Println("data dari channel", data)

	time.Sleep(4 * time.Second)
}

func SendParameter(channel chan string) {
	time.Sleep(2 * time.Second)
	// by default [chan] recieve data by reference not by value
	channel <- "this is parameter"
	fmt.Println("done donk!")
}

//  go test -v -run=TestChannelParameter
func TestChannelParameter(t *testing.T) {
	chn := make(chan string)
	defer close(chn)
	go SendParameter(chn)

	data := <-chn
	fmt.Println("data==>", data)

	time.Sleep(4 * time.Second)

}

func SendOnly(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "[[Input only]]"
	fmt.Println("done donk!")
}
func RecieveOnly(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println("[[output only]]", data)
}

//  go test -v -run=TestInOutChannel
func TestInOutChannel(t *testing.T) {
	chn := make(chan string)
	go SendOnly(chn)
	go RecieveOnly(chn)

	time.Sleep(3 * time.Second)
	close(chn)

}
