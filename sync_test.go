package main

import (
	"fmt"
	"sync"
	"testing"
)

var index = 0

func OnlyOnce() {
	index++
}

//  go test -v -run=TestWaitGroup
func TestSyncOnce(T *testing.T) {
	group := &sync.WaitGroup{}
	once := &sync.Once{}

	for i := 0; i < 30; i++ {
		// go RunAsynchronous(group, i)
		go func() {

			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()

	}
	// using group wait wil wait goroutine finish the process
	group.Wait()
	fmt.Println("[[[[[[[[[[[complete]]]]]]]", index)
}

var addToData = func(dt *sync.Map, val, key int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	dt.Store(val, key)
}

//  go test -v -run=TestMap
func TestMap(t *testing.T) {
	/**
	store(key,value)
	load(key)
	delete(key)
	range (func (key, value)
	*/

	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go addToData(data, i, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
