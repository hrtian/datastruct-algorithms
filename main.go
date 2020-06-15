package main

import (
	Circle_Queue "code.golang.com/Datastruct/04-Queue/3_Circle_Queue"
	"fmt"
)

func main() {
	q := Circle_Queue.New(4)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	fmt.Println(q)

	q.DeQueue()
	q.DeQueue()
	fmt.Println(q)
	q.EnQueue(0)
	q.EnQueue(-1)
	fmt.Println(q)
	q.EnQueue(5)
	fmt.Println(q)

	q.EnQueue(6)
	fmt.Println(q)

	q.EnQueue(7)
	q.EnQueue(8)
	q.EnQueue(9)
	q.EnQueue(10)

	fmt.Println(q)


}
