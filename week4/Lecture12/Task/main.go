package main

import (
	"fmt"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	var channel chan string = make(chan string,  bufferLimit)
	out := make(chan string, bufferLimit)
	go func ()  {
		for {
			if len(channel) == bufferLimit {
				 for i := 0; i < bufferLimit; i++ {
					 out <- <-channel
				 }
				 <-time.After(clearInterval)
			}

			channel <- data
			
		}
	}()
	
	return out
}

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		fmt.Println(f)
	}
}