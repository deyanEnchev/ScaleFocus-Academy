package main

import "fmt"

type Item struct {
	Value    int
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

// func toSlice(l *MagicList) []int {
// 	var mySlice []int
// 	mySlice = append([]int{l.LastItem.Value}, mySlice...)
// 	l.LastItem = l.LastItem.PrevItem
// 	if l.LastItem != nil {
// 		return append(toSlice(l), mySlice...)
// 	}
// 	return mySlice
// }

func toSlice(l *MagicList) []int {
	var currentItem = l.LastItem

	if currentItem.PrevItem != nil {
		var tmp []int
		l.LastItem = currentItem.PrevItem
		tmp = toSlice(l)
		return append(tmp, currentItem.Value)

	}

	return []int{currentItem.Value}
}

func main() {
	l := &MagicList{}
	add(l, 1)
	add(l, 2)
	add(l, 3)
	add(l, 4)


	var mySlice []int
	mySlice = toSlice(l)
	fmt.Println(mySlice)
}
