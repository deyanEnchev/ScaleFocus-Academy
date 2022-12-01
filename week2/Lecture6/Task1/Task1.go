package main

import "fmt"




type Item struct {
	Value int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func add(l *MagicList, value int) {

	item := Item{Value: value}

	if l.LastItem == nil {
		l.LastItem = &item
	} else {
		item.PrevItem = l.LastItem
		l.LastItem = &item
	}

}

func toSlice(l *MagicList) []int {
	var mySlice []int
	for l.LastItem != nil {
		mySlice = append([]int{l.LastItem.Value}, mySlice...)
		l.LastItem = l.LastItem.PrevItem
	
	}

	return mySlice
}

func main() {
	l := &MagicList{}
	add(l, 1)
	add(l, 2)
	add(l, 3)
	add(l, 4)
	add(l, 5)
	add(l, 6)
	add(l, 7)
	add(l, 8)
	add(l, 9)
	add(l, 10)

	fmt.Println(toSlice(l))
}