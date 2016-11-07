package main

import (
	"fmt"
	"./linkedlist"
)

func main() {
	l := linkedlist.New()
	l.Push("A")
	fmt.Println(l)
	l.Push("B")
	fmt.Println(l)
	l.Push(10)
	fmt.Println(l)
	l.Push(3.1415)
	fmt.Println(l)


	for it := l.Iterator(); linkedlist.HasNext(&it); {
		var e interface{} = linkedlist.Next(&it)
		fmt.Println(e)
	}

	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
}
