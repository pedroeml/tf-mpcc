package main

import (
	"fmt"
	"./linkedlist"
	"sync"
	"math/rand"
)

func main() {
	ch := make(chan interface{}, 10)
	var mutex = &sync.Mutex{}
	var wg = &sync.WaitGroup{}
	l := linkedlist.New()

	//l.Push("A")
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Push("B")
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Push(10)
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Push(3.1415)
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//
	//
	//for it := l.Iterator(); linkedlist.HasNext(&it); {
	//	var e interface{} = linkedlist.Next(&it)
	//	fmt.Println(e)
	//}
	//
	//l.Pop()
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Pop()
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Pop()
	//fmt.Println(l)
	//fmt.Printf("last = %s\n", l.Last().Value)
	//l.Pop()
	//fmt.Println(l)

	for i := 0; i < 3; i++ {
		l.Push(rand.Intn(10))
	}
	fmt.Println("LIST: ", l)
	testInserters(l, ch, mutex, wg)
}

func testInserters(l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	times := 5

	for i := 0; i < times; i++ {
		ch <- rand.Intn(10)
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		for i := 0; i < times; i++ {
			wg.Add(1)
			go insert(i, l, ch, mutex, wg)
		}

		for i := 0; i < times; i++ {
			wg.Add(1)
			go search(i, l, wg)
		}

		wg.Wait()

		for i := 0; i < times+3; i++ {
			wg.Add(1)
			go delete(i, l, ch, mutex, wg)
		}

		wg.Wait()
		wg2.Done()
	}()

	wg2.Wait()
}


func delete(deleterID int, l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println(">>> DELETER ", deleterID, " STARTED WORKING...")
	mutex.Lock()
	elem := l.Pop()
	if elem != nil {
		ch <- elem
		fmt.Println("#", deleterID, "DELETED:  ", elem, " LIST: ", l)
	}
	mutex.Unlock()
	wg.Done()
}

func insert(inserterID int, l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println(">>> INSERTER ", inserterID, " STARTED WORKING...")
	elem := <- ch
	mutex.Lock()
	l.Push(elem)
	fmt.Println("#", inserterID, "INSERTED: ", elem, " LIST: ", l)
	mutex.Unlock()
	wg.Done()
}

func search(searcherID int, l *linkedlist.SinglyLinkedList, wg *sync.WaitGroup) {
	fmt.Println(">>> SEARCHER ", searcherID," STARTED WORKING...")
	for it := l.Iterator(); linkedlist.HasNext(&it); {
		var e interface{} = linkedlist.Next(&it)
		fmt.Println("#", searcherID, "FOUND:   ", e)
	}
	wg.Done()
}
