package main

import (
	DEQueue "code.golang.com/Datastruct/04-Queue/2_DEQueue"
	"fmt"
)

func main() {
	q := DEQueue.New()
	q.EnQueueRear(22)
	q.EnQueueFront(11)
	q.EnQueueRear(33)
	q.EnQueueRear(44)

	fmt.Println(q)

	q.DeQueueFront()
	q.DeQueueRear()

	fmt.Println(q)

	for !q.IsEmpty() {
		fmt.Println(q.DeQueueFront())
		fmt.Println(q.IsEmpty())
	}
}
