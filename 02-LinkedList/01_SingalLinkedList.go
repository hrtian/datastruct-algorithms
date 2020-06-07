package main

import (
	"bytes"
	"fmt"
)

// 链表需要注重边界条件的测试 0 size size - 1

// ENotFound is a return-number for search element;
var ElementNotFound = -1

// SLL is a signal linked VNode;
type SLL struct {
	Size  int
	First *SNode
}

type SNode struct {
	Val  interface{}
	Next *SNode
}

func (l *SLL) clear() {
	l.Size = 0
	l.First = nil
}

func (l *SLL) size() int {
	return l.Size
}

func (l *SLL) isEmpty() bool {
	return l.Size == 0
}

func (l *SLL) contains(e interface{}) bool {
	return l.indexOf(e) != ElementNotFound
}

func (l *SLL) add(e interface{}) {
	l.addIndex(l.Size, e)
}

func (l *SLL) addIndex(index int, e interface{}) {
	if index == 0 {
		next := l.First
		l.First = &SNode{
			Val:  e,
			Next: next,
		}
	} else {
		prev := l.findNode(index - 1)
		next := prev.Next
		prev.Next = &SNode{
			Val:  e,
			Next: next,
		}
	}

	l.Size++
}

func (l *SLL) set(index int, e interface{}) interface{} {
	node := l.findNode(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *SLL) remove(index int) interface{} {
	old := l.First.Val
	if index == 0 {
		l.First = l.First.Next
	} else {
		prev := l.findNode(index - 1)
		old = prev.Next.Val
		prev.Next = prev.Next.Next
	}
	l.Size--
	return old
}

func (l *SLL) indexOf(e interface{}) int {
	node := l.First

	for i := 0; i < l.Size; i++ {
		if node.Val == e {
			return i
		}
		node = node.Next
	}
	return ElementNotFound
}

func (l *SLL) findNode(index int) *SNode {
	if index < 0 || index >= l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange", index, l.Size))
	}

	node := l.First

	for ; index > 0; index-- {
		node = node.Next
	}
	return node
}

func (l *SLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: SLL, size: %d, [", l.Size))
	node := l.First

	for i := 0; i < l.Size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(node.Val))
		node = node.Next
	}
	buffer.WriteString("]")
	return buffer.String()
}

func main() {
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

}
