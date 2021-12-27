package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) Enqueue(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) Dequeue() (Display string) {
	toDisplay := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toDisplay.data)
		toDisplay = toDisplay.next
		l.length--
	}
	return
}

func Display() {
	myList := linkedList{}
	node1 := node{data: 40}
	node2 := node{data: 20}
	node3 := node{data: 10}
	myList.Enqueue(&node1)
	myList.Enqueue(&node2)
	myList.Enqueue(&node3)
	myList.Dequeue()
}

func main() {
	Display()
}
