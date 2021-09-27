package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {

	var group sync.WaitGroup
	var counter int32 = 0
	var counterOld int32 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				counterOld++
				atomic.AddInt32(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter", counter, counterOld)
}
