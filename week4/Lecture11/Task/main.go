package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}

	evenChan := make(chan int)
	oddChan := make(chan int)

	for i := 0; i < len(inputs); i++ {
		wg.Add(2)

		//processEven
		go func(ch chan int, idx int) {
			number := inputs[idx]
			if number%2 == 0 {
				ch <- number
			}
			wg.Done()
		}(evenChan, i)
		
		//processOdd
		go func(ch chan int, idx int) {
			number := inputs[idx]
			if number%2 != 0 {
				ch <- number
			}
			wg.Done()
		}(oddChan, i)

		//Await for calls
		go func (eCh, oCh chan int) {
			select {
			case msg := <-eCh:
				fmt.Println(msg)
				<- oCh
			case msg2 := <-oCh:
				fmt.Println(msg2)
			}
		}(evenChan,oddChan)
		time.Sleep(1 * time.Nanosecond)
	}
}
