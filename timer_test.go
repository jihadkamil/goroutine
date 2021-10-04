package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// go test -v -run=TestTimer
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

// go test -v -run=TestTimeAfter
func TestTimeAfter(t *testing.T) {
	timer := time.After(3 * time.Second)

	fmt.Println(time.Now())

	time := <-timer
	fmt.Println(time)
}

// go test -v -run=TestTimeAfterFunc
func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {

		fmt.Println("time.Now(1)", time.Now())
		group.Done()

	})
	fmt.Println("time.Now(2)", time.Now())
	group.Wait()

}
