package main

import (
	"fmt"
	"strconv"
	"sync"
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

//  go test -v -run=TestBufferedChannel
func TestBufferedChannel(t *testing.T) {
	chn := make(chan string, 2)
	go func() {
		chn <- "first input"
		chn <- "second input"
		chn <- "third input" // third input will blocked / send error, because there are only 2 buffer unless one is empty
	}()

	go func() {

		fmt.Println("<-chn 1", <-chn)
		fmt.Println("<-chn 2", <-chn)
		// fmt.Println("<-chn 3", <-chn) // 3rd reciever  will blocked / send error, because there are only 2 buffer
	}()

	time.Sleep(4 * time.Second)
	fmt.Println("Seuleuseai")
}

//  go test -v -run=TestRangeChannel
func TestRangeChannel(t *testing.T) {
	chn := make(chan string)
	go func() {
		for index := 0; index < 10; index++ {
			chn <- "send " + strconv.Itoa(index+1)
		}
		close(chn) // close the "channel" is a must otherwise it will block the code/deadlock
	}()

	for data := range chn {
		fmt.Println("recieving data", data)
	}

	fmt.Println("Done donk!")
}

//  go test -v -run=TestSelectChannel
func TestSelectChannel(t *testing.T) {
	chn1 := make(chan string)
	chn2 := make(chan string)

	defer close(chn1)
	defer close(chn2)

	go SendOnly(chn1)
	go SendOnly(chn2)

	counter := 0
	for {
		select {
		case data := <-chn1:
			fmt.Println("dari chn 1", data)
			counter++
		case data := <-chn2:
			fmt.Println("dari chn 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

//  go test -v -run=TestMutex
func TestMutex(T *testing.T) {

	x := 0
	var mutex sync.Mutex
	go func() {
		for i := 0; i < 1000; i++ {
			go func() {
				for j := 0; j < 10; j++ {
					mutex.Lock() //hold until x++ run
					x++
					mutex.Unlock()

				}

			}()
		}
	}()
	time.Sleep(4 * time.Second)
	fmt.Println("ini hasil ", x)
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// send value sum to channel
	c <- sum
}

//  go test -v -run= TestSumArrayChannel
func TestSumArrayChannel(t *testing.T) {
	s := []int{18, 72, 23, 34, 52, 16, 71, 89}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("x+y", x+y)
}

// blocking channel
func TestChannelBlock(t *testing.T) {
	c := make(chan int)
	printed := "not be printed"
	/* func willbe blocked when data in channel not consumed by anyone
	go func() {
		printed = "be printed"
		<-c
	}()
	// */
	c <- 1
	fmt.Println("this code will", printed)
}

// buffered channel
