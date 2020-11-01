package Stack

import (
	"bytes"
	"code.golang.com/Datastruct/01-ArrayList"
	"fmt"
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

func (s *Stack) Peek() interface{} {
	return s.data.IndexOf(s.Size() - 1)
}

func (s *Stack) Clear() {
	s.data.Clear()
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Stack: ["))

	for i := 0; i < s.data.Size(); i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(s.data.Get(i)))
	}

	buffer.WriteString("]")

	return buffer.String()
}

//func main() {
//	s := New(6)
//	s.Push(0)
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	s.Push(4)
//	s.Push(5)
//	fmt.Println(s)
//	s.Pop()
//	fmt.Println(s)
//	fmt.Println(s.Top())
//}
