package _1_SLL

import (
	"bytes"
	"fmt"
)

// 链表需要注重边界条件的测试 0 size size - 1

// ENotFound is a return-number for search element;
var ElementNotFound = -1

// 01_SLL is a signal linked VNode;
type SLL struct {
	size  int
	first *SNode
}

type SNode struct {
	Val  interface{}
	Next *SNode
}

func (l *SLL) Clear() {
	l.size = 0
	l.first = nil
}

func (l *SLL) Size() int {
	return l.size
}

func (l *SLL) IsEmpty() bool {
	return l.size == 0
}

func (l *SLL) Contains(e interface{}) bool {
	return l.IndexOf(e) != ElementNotFound
}

func (l *SLL) Add(e interface{}) {
	l.AddIndex(l.size, e)
}

func (l *SLL) AddIndex(index int, e interface{}) {
	if index == 0 {
		next := l.first
		l.first = &SNode{
			Val:  e,
			Next: next,
		}
	} else {
		prev := l.Node(index - 1)
		next := prev.Next
		prev.Next = &SNode{
			Val:  e,
			Next: next,
		}
	}

	l.size++
}

func (l *SLL) Set(index int, e interface{}) interface{} {
	node := l.Node(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *SLL) Remove(index int) interface{} {
	old := l.first.Val
	if index == 0 {
		l.first = l.first.Next
	} else {
		prev := l.Node(index - 1)
		old = prev.Next.Val
		prev.Next = prev.Next.Next
	}
	l.size--
	return old
}

func (l *SLL) IndexOf(e interface{}) int {
	node := l.first

	for i := 0; i < l.size; i++ {
		if node.Val == e {
			return i
		}
		node = node.Next
	}
	return ElementNotFound
}

func (l *SLL) Node(index int) *SNode {
	if index < 0 || index >= l.size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange", index, l.size))
	}

	node := l.first

	for ; index > 0; index-- {
		node = node.Next
	}
	return node
}

func (l *SLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: 01_SLL, size: %d, [", l.size))
	node := l.first

	for i := 0; i < l.size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(node.Val))
		node = node.Next
	}
	buffer.WriteString("]")
	return buffer.String()
}

//func main() {
//_ := &SNode{
//	Val: 1,
//	Next: &SNode{
//		Val: 4,
//		Next: &SNode{
//			Val: 2,
//			Next: &SNode{
//				Val: 5,
//				Next: &SNode{
//					Val: 3,
//					Next: &SNode{
//						Val:  6,
//						Next: nil,
//					},
//				},
//			},
//		},
//	},
//}
//}
