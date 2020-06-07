package main

import (
	"code.golang.com/Datastruct/03-Stack/01_ArrayStack"
	"fmt"
)

func main() {
	s := Stack.New(6)
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	for s.IsEmpty() == false {
		fmt.Println(s.Pop())
	}

	fmt.Println(s)
}
