package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(goPrimesAndSleep(50, 1*time.Microsecond))
	fmt.Println(primesAndSleep(50, 1*time.Microsecond))
}

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	var wg sync.WaitGroup
	res := []int{}
	ch := make(chan int)
	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(c chan<- int, k int) {
			for i := 2; i < n; i++ {
				if k%i == 0 {
					defer wg.Done()
					time.Sleep(sleep)
					if k == i {
						c <- k
					}
					break
				}
			}
		}(ch, k)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for item := range ch {
		res = append(res, item)
	}

	return res
}

func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}
