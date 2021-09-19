package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pool => saving data then u can use it by accessing pool

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// default value
		New: func() interface{} {
			return "Halo default"
		},
	}
	group := &sync.WaitGroup{}

	pool.Put("Heyy 1")
	pool.Put("halo 2")
	pool.Put("Heyy 3")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println("data get from pool", data)
			// time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	time.Sleep(11 * time.Second)
	fmt.Println("[[[[selecseai]]]]")

}
