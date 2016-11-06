package linkedlist

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	next  *Node
	list  *SinglyLinkedList
	Value interface{}
}

func (n *Node) Next() *Node {
	if n.list != nil && n.next != nil {
		return n.next
	}
	return nil
}

type SinglyLinkedList struct {
	head Node
	size int
}

// Initializes or clears list l.
func (l *SinglyLinkedList) Init() *SinglyLinkedList {
	l.head.next = nil
	l.size = 0
	return l
}

// Returns an initialized list.
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}

// Returns the number of nodes of list l.
func (l *SinglyLinkedList) Size() int {
	return l.size
}

// Returns the first Node of list l or nil if list is empty.
func (l *SinglyLinkedList) First() *Node {
	if l.size == 0 {
		return nil
	}
	return l.head.next
}

// Returns the last Node of list l or nil if list is empty.
func (l *SinglyLinkedList) Last() *Node {
	if l.size == 0 {
		return nil
	}

	var n *Node

	for n = l.head.next; n.next != nil; n = n.next {

	}

	return n
}

// Inserts a new node next to the node at, increments l.size and returns the inserted node.
func (l *SinglyLinkedList) insert(node, at *Node) *Node {
	var n = at.next
	at.next = node
	node.next = n
	node.list = l
	l.size++
	return node
}

// Insert a new node next to the node at and returns the inserted node.
func (l *SinglyLinkedList) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

// Inserts a new node in the end of list l.
func (l *SinglyLinkedList) Push(v interface{}) {
	var last = l.Last()
	if last == nil {
		l.insertValue(v, &l.head)
	} else {
		l.insertValue(v, last)
	}
}

// Removes node from its list, decrements l.size and returns the removed node.
func (l *SinglyLinkedList) remove(node *Node) *Node {
	var previous *Node
	if l.head.next != nil {
		for previous = l.head.next; previous.next != nil && previous.next != node; previous = previous.next {

		}
	}

	previous.next = node.next
	node.next = nil // avoid memory leaks
	node.list = nil
	l.size--
	return node
}

// Removes node from list l and returns it's value.
func (l *SinglyLinkedList) Remove(node *Node) interface{} {
	if node.list == l {
		l.remove(node)
	}
	return node.Value
}

// Removes the list l's last node and returns it's value.
func (l *SinglyLinkedList) Pop() interface{} {
	return l.Remove(l.Last())
}

// Returns a toString like of list l.
func (l *SinglyLinkedList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	if l.size != 0 && l.head.next != nil {
		for node := l.head.next; node != nil; node = node.next {
			switch v := node.Value.(type) {
			case string:
				buffer.WriteString(v)
			case int:
				buffer.WriteString(strconv.Itoa(v))
			case float32:
			case float64:
				buffer.WriteString(fmt.Sprintf("%.2f", v))
			default:
				fmt.Println("Node's value isn't a string, int nor a float")
			}

			if node.next != nil {
				buffer.WriteString(", ")
			}
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}
