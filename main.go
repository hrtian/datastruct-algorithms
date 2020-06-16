package main

import (
	Circle_DEQueue "code.golang.com/Datastruct/04-Queue/4_Circle_DEQueue"
	"fmt"
)

func main() {
	q := Circle_DEQueue.New(4)
	q.EnQueueRear(2)
	q.EnQueueFront(1)
	q.EnQueueRear(3)
	q.EnQueueRear(4)
	fmt.Println(q) // 1 2 3 4

	q.DeQueueFront()
	q.DeQueueRear()
	fmt.Println(q) // _ 2 3 _

	q.EnQueueFront(1)
	q.EnQueueFront(4)
	fmt.Println(q)

	q.EnQueueFront(0)
	q.EnQueueFront(0)
	q.EnQueueFront(0)
	q.EnQueueFront(0)
	fmt.Println(q)
}
