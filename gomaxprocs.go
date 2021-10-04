package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	group := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		group.Add(1)

		go func() {
			time.Sleep(3 * time.Second)
			group.Done()

		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("totalCpu", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("totalThread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("totalGoRoutine", totalGoRoutine)

	group.Wait()

}
