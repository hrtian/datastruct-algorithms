package dll

import (
	"bytes"
	"fmt"
)

// DNode is double linked list node
type DNode struct {
	Val  interface{}
	Prev *DNode
	Next *DNode
}

// 03_DLL is a struct for double linked list
type DLL struct {
	size  int
	first *DNode
	last  *DNode
}

func New() *DLL {
	list := new(DLL)

	list.first = new(DNode)
	list.last = new(DNode)
	list.first.Next = list.last
	list.last.Prev = list.first

	return list
}

func (l *DLL) Size() int {
	return l.size
}

func (l *DLL) IsEmpty() bool {
	return l.size == 0
}

func (l *DLL) AddIndex(index int, e interface{}) {
	if index > l.size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.size))
	}

	var aim *DNode
	if index == 0 {
		aim = l.first.Next
	} else if index == l.size {
		aim = l.last
	} else {
		aim = l.Node(index)
	}

	prev := aim.Prev
	newNode := &DNode{
		Val:  e,
		Prev: prev,
		Next: aim,
	}

	aim.Prev = newNode
	prev.Next = newNode
	l.size++
}

func (l *DLL) Add(e interface{}) {
	l.AddIndex(l.size, e)
}

func (l *DLL) Set(index int, e interface{}) interface{} {
	node := l.Node(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *DLL) Remove(index int) interface{} {
	aim := l.Node(index)
	prev := aim.Prev
	next := aim.Next
	old := aim.Val

	next.Prev = prev
	prev.Next = next
	l.size--
	return old
}

func (l *DLL) Get(index int) interface{} {
	return l.Node(index).Val
}

func (l *DLL) IndexOf(e interface{}) int {
	if l.size == 0 {
		return -1
	}

	node := l.first.Next
	for i := 0; i < l.size; i++ {
		if node.Val == e {
			return i
		}

		node = node.Next
	}

	return -1
}

func (l *DLL) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

func (l *DLL) Node(index int) *DNode {
	if index < 0 || index >= l.size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.size))
	}

	var node *DNode
	if index < l.size>>1 {
		node = l.first.Next
		for i := 0; i < index; i++ {
			node = node.Next
		}
	} else {
		node = l.last.Prev
		for i := l.size - 1; i > index; i-- {
			node = node.Prev
		}
	}

	return node
}

func (l *DLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: 03_DLL, size: %d, [", l.size))
	node := l.first.Next

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
//	list := newDLL()
//	list.add(0)
//	fmt.Println(list.size, list)
//	list.add(1)
//	fmt.Println(list.size, list)
//	list.add(2)
//	fmt.Println(list.size, list)
//	list.add(3)
//	fmt.Println(list.size, list)
//}
