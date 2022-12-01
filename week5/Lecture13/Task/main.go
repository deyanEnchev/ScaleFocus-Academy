package main

import (
	"context"
	"log"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer     chan string
	cancelFunc context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	bc := BufferedContext{Context: ctx,
		buffer:     make(chan string, bufferSize),
		cancelFunc: cancel}

	return &bc
}

func (bc *BufferedContext) Done() <-chan struct{} {

	if len(bc.buffer) == cap(bc.buffer) {
		bc.cancelFunc()
	}

	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc, bc.buffer)
}

func main() {
	ctx := NewBufferedContext(10*time.Second, 50) // <- from here?
	//From where does the context.WithTimeout start counting?
	ctx.Run(func(ctx context.Context, buffer chan string) { // <- or from here
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 500) // try different values here
				log.Println("bar")
			}
		}
	})
}
