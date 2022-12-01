package main

import (
	"errors"
	"fmt"
	"log"
)

type Action func() error

func SafeExecWithError(a Action) Action {
	defer func() error {
		if r := recover(); r != nil {
			return fmt.Errorf("safe exec: %w", r)
		}
		return fmt.Errorf("couldn't recover")
	}()

	a = func() error {
		return errors.New("wrapped error")
	}

	return a
}

func SafeExecWithPanic(a Action) Action {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered successfully")
		} else {
			fmt.Println("couldn't recover")
		}
	}()

	a = func() error {
		return errors.New("wrapped error")
	}

	panic("boom")

	return a
}

func SafeExecWithoutAnError(a Action) Action {
	defer func() error {
		return nil
	}()

	return nil
}

func main() {
	var a Action

	//Try each one by commenting the others.
	
	// 1. With an Error

	err := SafeExecWithError(a)
	if err != nil {
		log.Fatalf("there was an error: %v", err())
	} else {
		fmt.Println("no error!")
	}

	// 2. With a panic

	err2 := SafeExecWithPanic(a)
	if err2 != nil {
		log.Fatalf("there was an error2: %v", err2())
	} else {
		fmt.Println("no error!")
	}

	// 3. Without error and panic

	err3 := SafeExecWithoutAnError(a)
	if err3 != nil {
		log.Fatalf("there was an error3: %v", err3())
	} else {
		fmt.Println("no error!")
	}
}
