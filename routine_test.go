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

func HelloWow() {
	fmt.Println("hello worwld")
}

// /*
func TestHelloWow(T *testing.T) {

	go HelloWow()
	fmt.Println("berhasil")
	time.Sleep(1 * time.Second)
}

// */
func DisplayNumber(i int) {
	fmt.Println("display", i)
}
func TestManyGoroutine(T *testing.T) {
	for i := 0; i < 1000; i++ {
		// test: how light goroutine is
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
