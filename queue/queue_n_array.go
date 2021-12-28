package main

import "fmt"

func enqueue(q []int, num int) (result []int) {
    q = append(q, num)
    return q
}

func dequeue(q []int) ([]int, int) {
    return q[1:], q[0]
}

func main() {
    var queue []int

    // enqueue elements to queue
    // 1 oldest -> 10 newest
    for i := 1; i <= 10; i++ {
        queue = enqueue(queue, i)
    }

    // dequeue current oldest element : 1
    queue, element := dequeue(queue)
    fmt.Println("current queue after 1st dequeue: ", queue)
    fmt.Println("dequeued element: ", element)
   
    // dequeue current oldest element : 2
    queue, element = dequeue(queue)
    fmt.Println("current queue after 2nd dequeue: ", queue)
    fmt.Println("dequeued item: ", element)
}