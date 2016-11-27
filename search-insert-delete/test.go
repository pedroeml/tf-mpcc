package main

import (
	"fmt"
	"./linkedlist"
	"sync"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan interface{}, 20)
	var mutex = &sync.Mutex{}
	var wg = &sync.WaitGroup{}
	l := linkedlist.New()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 5; i++ {
		l.Push(rand.Intn(10))
	}
	fmt.Println("LIST: ", l)

	//testSearchers(l, wg)
	//testInserters(l, ch, mutex, wg)
	//testDeleters(l, ch, mutex, wg)
	//testSearchersInserters(l, ch, mutex, wg)
	testSearchInsertDelete(l, ch, mutex, wg)
}

/** This test proves that searchers can run concurrently with only one inserter,
 * only one inserter executes each time at the list and deleter runs in mutual exclusion:
 * one deleter runs at the time and no other inserter or searcher must be running.
 */
func testSearchInsertDelete(l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("\nSEARCH-INSERT-DELETE\n")
	times := 5

	for i := 0; i < times*2; i++ {
		ch <- rand.Intn(10)
	}

	searcherExecutedOnce := false
	inserterExecutedOnce := false
	deleterExecutedOnce := false

	for i := 0; !(searcherExecutedOnce && inserterExecutedOnce && deleterExecutedOnce) && i < times*2; i++ {	// this condition requires that all sort of thread (searcher, inserter and deleter) need to run at least once.
		switch rand.Intn(3) {
		case 0:
			searcherExecutedOnce = true
			for j := 0; j < rand.Intn(times-2) + 1; j++ {
				wg.Add(1)
				go search((i+j)*(j+3), l, wg)
			}
		case 1:
			inserterExecutedOnce = true
			for j := 0; j < rand.Intn(times-2) + 1; j++ {
				wg.Add(1)
				go insert((i+j)*(j+7), l, ch, mutex, wg)
			}
		case 2:
			deleterExecutedOnce = true
			wg.Wait()	// Rarely it might generates a fatal error: all goroutines are asleep. It only happens if there is no other goroutine running. At the moment, no idea how to check if there are any goroutine running to perform Wait().
			for j := 0; j < rand.Intn(times-2) + 1; j++ {
				wg.Add(1)
				go delete((i+j)*(j+11), l, ch, mutex, wg)
			}
			wg.Wait()
		}
	}

	wg.Wait()
}

// This test proves that searchers can run concurrently with only one inserter.
func testSearchersInserters(l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("\nSEARCH-INSERT\n")
	times := 5

	for i := 0; i < times; i++ {
		ch <- rand.Intn(10)
	}

	for i := 0; i < times; i++ {
		wg.Add(1)
		go insert(i, l, ch, mutex, wg)
	}

	for i := 0; i < times; i++ {
		wg.Add(1)
		go search(i, l, wg)
	}

	wg.Wait()
}

// This test proves that searchers can run concurrently with others searchers.
func testSearchers(l *linkedlist.SinglyLinkedList, wg *sync.WaitGroup) {
	fmt.Println("\nSEARCH\n")
	times := 5

	for i := 0; i < times; i++ {
		wg.Add(1)
		go search(i, l, wg)
	}

	wg.Wait()
}

// This test proves that only one inserter executes each time.
func testInserters(l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("\nINSERT\n")
	times := 5

	for i := 0; i < times; i++ {
		ch <- rand.Intn(10)
	}

	for i := 0; i < times; i++ {
		wg.Add(1)
		go insert(i, l, ch, mutex, wg)
	}

	wg.Wait()
}

// This test proves that only one deleter executes each time.
func testDeleters(l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("\nDELETE\n")
	times := 5

	for i := 0; i < times; i++ {
		wg.Add(1)
		go delete(i, l, ch, mutex, wg)
	}

	wg.Wait()

}

func delete(deleterID int, l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	if l.Size() > 0 {
		fmt.Println(">>> DELETER ", deleterID, " STARTED WORKING...")
		elem := l.Pop()
		ch <- elem
		fmt.Println("#", deleterID, "DELETED:  ", elem, " LIST: ", l)
	}
	mutex.Unlock()
	wg.Done()
}

func insert(inserterID int, l *linkedlist.SinglyLinkedList, ch chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println(">>> INSERTER ", inserterID, " STARTED WORKING...")
	elem := <- ch
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
	fmt.Println("--- SEARCHER ", searcherID," FINISHED WORKING...")
	wg.Done()
}
