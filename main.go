package main

import (
	Queue "code.golang.com/Datastruct/04-Queue"
	"fmt"
)

func main(){
	q := Queue.New()
	q.EnQueue(11)
	q.EnQueue(22)
	q.EnQueue(33)
	q.EnQueue(44)

	for !q.IsEmpty() {
		fmt.Println(q.DeQueue())
		fmt.Println(q.IsEmpty())
	}

}