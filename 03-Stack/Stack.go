package mian

import "fmt"

type stack struct {
	Size int
	Val  []interface{}
}

func (s *stack) size() int {
	return s.Size
}

func (s *stack) isEmpty() bool {
	return s.Size != 0
}

func (s *stack) push(e interface{}) {
	s.Val[s.Size] = e
	s.Size++
}

func (s *stack) pop() interface{} {
	if s.Size == 0 {
		panic(fmt.Sprintf("Out Of Range, size: %d\n", s.Size))
	}
	s.Size--
	return s.Val[s.Size-1]
}

func (s *stack) top() interface{} {
	return s.Val[s.Size-1]
}
