package main

import (
	"fmt"
	"os"
	"time"
)

func Timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func Watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\n time out! no answer more thane ", timeout, " seconds")
	os.Exit(0)
}
func main() {
	timeout := 5
	ch := make(chan bool)

	go Timer(timeout, ch)
	go Watcher(timeout, ch)

	var input string

	fmt.Println("1 + 1 ?")
	fmt.Scan(&input)

	if input == "2" {
		fmt.Println("cool! the answer is right (not left)!!, \n you look very smart today :)")
	} else {
		fmt.Println("kumaha sih!")
	}

}
