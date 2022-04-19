package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//  go test -v -run=TestAtomic
func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var counter int32 = 0
	var counterOld int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 10; j++ {
				counterOld++
				// use atomic only for primitive data structure, use Mutex for struct
				atomic.AddInt32(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter", counter, counterOld)
}

//  go test -v -run=TestRaceCondition
func TestRaceCondition(t *testing.T) {
	// /*
	counterOld := 0

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				counterOld = counterOld + 1
				// use atomic only for primitive data structure, use Mutex for struct
			}
		}()
	}
	// */

	time.Sleep(4 * time.Second)
	fmt.Println("counter Race condition", counterOld)
}

//  go test -v -run=TestRaceCondition
func TestMutexRaceCondition(t *testing.T) {
	// /*
	counterOld := 0
	var mut sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				mut.Lock()
				counterOld++
				mut.Unlock()
				// use atomic only for primitive data structure, use Mutex for struct
			}
		}()
	}
	// */

	time.Sleep(5 * time.Second)
	fmt.Println("counter Race condition", counterOld)
}

type BankAccount struct {
	Balance int
	RWMutex sync.RWMutex
}

func (acc *BankAccount) GetBalance() int {
	acc.RWMutex.RLock()
	balance := acc.Balance
	acc.RWMutex.RUnlock()
	return balance

}

func (acc *BankAccount) AddBalance(amount int) {
	acc.RWMutex.Lock()
	acc.Balance = acc.Balance + amount
	acc.RWMutex.Unlock()
}

func TestRWMutex(T *testing.T) {
	acc := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				acc.AddBalance(1)
				fmt.Println(acc.GetBalance())

			}
		}()

	}
	time.Sleep(4 * time.Second)
}

func TestSR(T *testing.T) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
