package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var group = sync.WaitGroup{}
var cond = sync.NewCond(&locker)

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	// locking condition
	cond.L.Lock()
	// wait a moment
	cond.Wait()
	fmt.Println("[[[[[Done]]]]]", value)
	cond.L.Unlock()
}

//  go test -v -run=TestInOutChannel
func TestWaitCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		// group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			// send signal to run 1 [one] goroutine, the rest will wait
			cond.Signal()

		}
	}()

	/*
		go func(){
			time.Sleep(1 *time.Second)
			// will run all goroutine automatically
			cond.Broadcast()
		}
	*/

	group.Wait()

}
