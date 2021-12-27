package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()

	// enqueue elements to queue
	// 1 oldest -> 10 newest
	for i := 1; i <= 10; i++ {
		queue.PushBack(i)
	}

	// for every elements in the queue
	for queue.Len() > 0 {
		// take the oldest element
		element := queue.Front()
		fmt.Println("dequeued element: ", element.Value)

		// remove the oldest element
		// so it's replaced by
		// the second oldest element
		queue.Remove(element)
	}
}
