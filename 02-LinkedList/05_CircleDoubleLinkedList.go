package main

import (
	"bytes"
	"fmt"
)

// CDNode is double linked list node
type CDNode struct {
	Val  interface{}
	Prev *CDNode
	Next *CDNode
}

// DLL is a struct for double linked list
type CDLL struct {
	Size  int
	First *CDNode
	Last  *CDNode
}

func newCDLL() *CDLL {
	list := new(CDLL)

	list.First = new(CDNode)
	list.Last = new(CDNode)
	list.First.Next = list.Last
	list.Last.Prev = list.First

	return list
}

func (l *CDLL) addIndex(index int, e interface{}) {
	if index > l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.Size))
	}

	var aim *CDNode
	if index == 0 {
		aim = l.First.Next
	} else if index == l.Size {
		aim = l.Last
	} else {
		aim = l.node(index)
	}

	prev := aim.Prev
	newNode := &CDNode{
		Val:  e,
		Prev: prev,
		Next: aim,
	}

	aim.Prev = newNode
	prev.Next = newNode
	l.Size++
}

func (l *CDLL) add(e interface{}) {
	l.addIndex(l.Size, e)
}

func (l *CDLL) set(index int, e interface{}) interface{} {
	node := l.node(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *CDLL) remove(index int) interface{} {
	aim := l.node(index)
	prev := aim.Prev
	next := aim.Next
	old := aim.Val

	next.Prev = prev
	prev.Next = next

	return old
}

func (l *CDLL) get(index int) interface{} {
	return l.node(index).Val
}

func (l *CDLL) indexOf(e interface{}) int {
	if l.Size == 0 {
		return -1
	}

	node := l.First.Next
	for i := 0; i < l.Size; i++ {
		if node.Val == e {
			return i
		}

		node = node.Next
	}

	return -1
}

func (l *CDLL) clear() {
	l.Size = 0
	l.First = nil
	l.Last = nil
}

func (l *CDLL) node(index int) *CDNode {
	if index < 0 || index >= l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.Size))
	}

	var node *CDNode
	if index < l.Size>>1 {
		node = l.First.Next
		for i := 0; i < index; i++ {
			node = node.Next
		}
	} else {
		node = l.Last.Prev
		for i := l.Size - 1; i > index; i-- {
			node = node.Prev
		}
	}

	return node
}

func (l *CDLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: CDLL, size: %d, [", l.Size))
	node := l.First.Next

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
	list := newCDLL()
	fmt.Println(list)
}
