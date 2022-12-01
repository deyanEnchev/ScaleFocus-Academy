package main

import (
	"fmt"
	"testing"
	"time"
)

var num = 100
var table = []struct {
	sleep time.Duration
}{
	{sleep: 0 * time.Millisecond},
	{sleep: 5 * time.Millisecond},
	{sleep: 10 * time.Millisecond},
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_time_%v", v.sleep), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				primesAndSleep(num, v.sleep)
			}
		})
	}
}

func BenchmarkGoPrimesAndSleep(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_time_%v", v.sleep), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				goPrimesAndSleep(num, v.sleep)
			}
		})
	}
}
