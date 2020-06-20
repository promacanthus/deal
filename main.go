package main

import (
	"deal/linkedlist"
)

func main() {
	list := linkedlist.NewSingleLinkedList()
	for i := 1; i <= 10; i++ {
		list.Last(i)
	}
	list.Print()
	list.RemoveLastN(3)
	list.Print()
}
