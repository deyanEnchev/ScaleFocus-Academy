package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func GroupBy[T any, U comparable](orders []T, keyFn func(T) U) map[U][]T {
	m := make(map[U][]T, len(orders))

	for _, o := range orders {
		cust := keyFn(o)
		m[cust] = append(m[cust], o)
	}

	return m
}

func main() {
	orders := []Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}

	results := GroupBy(orders, func(o Order) string { return o.Customer})

	fmt.Print(results)
}