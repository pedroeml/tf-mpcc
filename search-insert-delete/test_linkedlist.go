package main

import (
	"fmt"
	"./linkedlist"
)

func main() {
	l := linkedlist.New()
	l.Push("A")
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Push("B")
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Push(10)
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Push(3.1415)
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)


	for it := l.Iterator(); linkedlist.HasNext(&it); {
		var e interface{} = linkedlist.Next(&it)
		fmt.Println(e)
	}

	l.Pop()
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Pop()
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Pop()
	fmt.Println(l)
	fmt.Printf("last = %s\n", l.Last().Value)
	l.Pop()
	fmt.Println(l)
}
