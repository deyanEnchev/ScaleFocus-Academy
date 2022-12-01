package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	counter  int
	lastItem string
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {

	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()

			if cp.counter == times {
				cp.Unlock()
				break
			}

			if cp.lastItem != "foo" {
				fmt.Print("foo")
				cp.lastItem = "foo"
				cp.counter++
			}

			cp.Unlock()
		}
	}()

}
func (cp *ConcurrentPrinter) PrintBar(times int) {

	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()

			if cp.counter == times {
				cp.Unlock()
				break
			}

			if cp.lastItem != "bar" {
				fmt.Print("bar")
				cp.lastItem = "bar"
				cp.counter++
			}

			cp.Unlock()
		}
	}()

}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.lastItem = "bar"
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()
}
