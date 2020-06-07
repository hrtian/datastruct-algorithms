package stack

import (
	"code.golang.com/Datastruct/01-ArrayList"
)

type stack struct {
	Val *arrayList.ArrayList
}

func New(capacity int) *stack{
	return &stack{
		Val: arrayList.NewArrayList(capacity),
	}
}

func (s *stack) size() int {
	return s.Val.Size
}

func (s *stack) isEmpty() bool {
}

func (s *stack) push(e interface{}) {
}

func (s *stack) pop() interface{} {
}

func (s *stack) top() interface{} {
}
