package main

import (
	"fmt"
	"testing"
	"time"
)

/*
 run : {
	 go test -v
	 go test -v -run=TestManyGoroutine
 }
//  */

func HelloWow(text string) {
	fmt.Println("hello worwld", text)
}

// /*
func TestHelloWow(T *testing.T) {

	go HelloWow("1")
	go HelloWow("2")
	go HelloWow("3")
	go HelloWow("4")
	go HelloWow("5")
	go HelloWow("6")
	go HelloWow("7")
	fmt.Println("berhasil")
	time.Sleep(1 * time.Second)
}

// */
func DisplayNumber(i int) {
	fmt.Println("display", i)

}
func TestManyGoroutine(T *testing.T) {
	for i := 0; i <= 100000; i++ {
		// test: how light goroutine is
		go DisplayNumber(i)
	}
	fmt.Println("done")
	time.Sleep(5 * time.Second)
}

func TestCreateChannelLagi(t *testing.T) {
	newChan := make(chan string)

	defer close(newChan)

	go func() {
		time.Sleep(2 * time.Second)
		newChan <- "jihad kamil"
		fmt.Println("channel already assigned")

	}()

	data := <-newChan
	fmt.Println("nih data===>", data)
	time.Sleep(4 * time.Second)
}
