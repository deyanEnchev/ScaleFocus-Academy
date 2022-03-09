package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

// type PrintText interface {
// 	Print(string)
// }

// func (cp *ConcurrentPrinter) Print(text string) {
// 	cp.Lock()
// 	fmt.Print(text)
// 	cp.Unlock()
// }

func (cp *ConcurrentPrinter) PrintFoo(times int) {

	go func() {
		for i := 0; i < times; i++ {
			defer cp.Done()
			cp.Lock()
			cp.Add(1)
			if i%2 == 0 {
				fmt.Print("foo")
			}
			cp.Unlock()
			time.Sleep(time.Millisecond)
		}
	}()

}
func (cp *ConcurrentPrinter) PrintBar(times int) {

	go func() {
		for i := 0; i < times; i++ {
			defer cp.Done()
			cp.Lock()
			cp.Add(1)
			if i%2 != 0 {
				fmt.Print("bar")
			}
			cp.Unlock()

			time.Sleep(time.Millisecond)
		}
	}()


}

func main() {
	times := 10

	cp := &ConcurrentPrinter{}
	cp.PrintFoo(times)
	cp.PrintBar(times)
	time.Sleep(time.Millisecond)
	cp.Wait()
}
