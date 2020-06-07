package main

import(
	"code.golang.com/Datastruct/03-Stack/01_ArrayStack"
	"fmt"
)

func main() {
	s := Stack.New(6)
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.Top())
}