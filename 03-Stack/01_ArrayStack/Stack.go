package Stack

import (
	"code.golang.com/Datastruct/01-ArrayList"
)

type Stack struct {
	data *arrayList.ArrayList
}

func New(capacity int) *Stack {
	return &Stack{
		data: arrayList.New(capacity),
	}
}

func (s *Stack) Size() int {
	return s.data.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s *Stack) Push(e interface{}) {
	s.data.Add(e)
}

func (s *Stack) Pop() interface{} {
	return s.data.RemoveIndex(s.Size() - 1)
}

func (s *Stack) Top() interface{} {
	return s.data.IndexOf(s.Size() - 1)
}
