package csll

import (
	"bytes"
	"fmt"
)

type CSNode struct {
	Val  interface{}
	Next *CSNode
}

type CLL struct {
	Size  int
	First *CSNode
}

func (l *CLL) clear() {
	l.Size = 0
	l.First = nil
}

func (l *CLL) size() int {
	return l.Size
}

func (l *CLL) isEmpty() bool {
	return l.Size == 0
}

func (l *CLL) contains(e interface{}) bool {
	return l.indexOf(e) != -1
}

func (l *CLL) add(e interface{}) {
	l.addIndex(l.Size, e)
}

func (l *CLL) addIndex(index int, e interface{}) {
	if index == 0 {
		var cNode *CSNode
		if l.Size == 0 {
			cNode = &CSNode{
				Val:  e,
				Next: nil,
			}
			cNode.Next = cNode
		} else {
			last := l.node(l.Size - 1)
			cNode = &CSNode{
				Val:  e,
				Next: l.First,
			}
			last.Next = cNode
		}

		l.First = cNode
	} else {
		prev := l.node(index - 1)
		next := prev.Next
		prev.Next = &CSNode{
			Val:  e,
			Next: next,
		}
	}

	l.Size++
}

func (l *CLL) set(index int, e interface{}) interface{} {
	node := l.node(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *CLL) remove(index int) interface{} {
	old := l.First.Val

	if index == 0 {
		if l.Size == 1 {
			l.First = nil
		} else {
			l.First = l.First.Next
			last := l.node(l.Size - 1)
			last.Next = l.First
		}
	} else {
		prev := l.node(index - 1)
		old = prev.Next.Val
		prev.Next = prev.Next.Next
	}

	l.Size--
	return old
}

func (l *CLL) indexOf(e interface{}) int {
	node := l.First

	for i := 0; i < l.Size; i++ {
		if node.Val == e {
			return i
		}
		node = node.Next
	}
	return -1
}

func (l *CLL) node(index int) *CSNode {
	if index < 0 || index >= l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange", index, l.Size))
	}

	node := l.First

	for ; index > 0; index-- {
		node = node.Next
	}
	return node
}

func (l *CLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: CLL, size: %d, [", l.Size))
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
	list := new(CLL)
	list.add(0)
	list.add(1)
	list.add(2)
	list.add(3)
	list.remove(0)
	fmt.Println(list.node(2).Next.Val)
	list.remove(2)
	fmt.Println(list.node(1).Next.Val)
	fmt.Println(list)
}
