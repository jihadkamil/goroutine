package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, index int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("berhasil-------", index)
	time.Sleep(1 * time.Second)
}

//  go test -v -run=TestWaitGroup
func TestWaitGroup(T *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		fmt.Println("index", i)
		go RunAsynchronous(group, i)
	}
	// using group wait wil wait goroutine finish the process
	group.Wait()

	fmt.Println("[[[[[[[[[[[complete]]]]]]]")
}
