package main

import (
	"reflect"
	"testing"
)

func TestOrderBy(t *testing.T) {
	orders := []Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}

	got := GroupBy(orders, func(o Order) string { return o.Customer })
	want := map[string][]Order{
		"John": {{"John", 1000}, {"John", 1200}},
		"Sara": {{"Sara", 2000}, {"Sara", 1800}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Test failed, wanted %v, got %v", want, got)
	}
}
