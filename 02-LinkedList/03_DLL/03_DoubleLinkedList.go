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
	Size  int
	First *DNode
	Last  *DNode
}

func newDLL() *DLL {
	list := new(DLL)

	list.First = new(DNode)
	list.Last = new(DNode)
	list.First.Next = list.Last
	list.Last.Prev = list.First

	return list
}

func (l *DLL) addIndex(index int, e interface{}) {
	if index > l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.Size))
	}

	var aim *DNode
	if index == 0 {
		aim = l.First.Next
	} else if index == l.Size {
		aim = l.Last
	} else {
		aim = l.node(index)
	}

	prev := aim.Prev
	newNode := &DNode{
		Val:  e,
		Prev: prev,
		Next: aim,
	}

	aim.Prev = newNode
	prev.Next = newNode
	l.Size++
}

func (l *DLL) add(e interface{}) {
	l.addIndex(l.Size, e)
}

func (l *DLL) set(index int, e interface{}) interface{} {
	node := l.node(index)
	old := node.Val
	node.Val = e
	return old
}

func (l *DLL) remove(index int) interface{} {
	aim := l.node(index)
	prev := aim.Prev
	next := aim.Next
	old := aim.Val

	next.Prev = prev
	prev.Next = next

	return old
}

func (l *DLL) get(index int) interface{} {
	return l.node(index).Val
}

func (l *DLL) indexOf(e interface{}) int {
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

func (l *DLL) clear() {
	l.Size = 0
	l.First = nil
	l.Last = nil
}

func (l *DLL) node(index int) *DNode {
	if index < 0 || index >= l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange\n", index, l.Size))
	}

	var node *DNode
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

func (l *DLL) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: 03_DLL, size: %d, [", l.Size))
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

//func main() {
//	list := newDLL()
//	list.add(0)
//	fmt.Println(list.Size, list)
//	list.add(1)
//	fmt.Println(list.Size, list)
//	list.add(2)
//	fmt.Println(list.Size, list)
//	list.add(3)
//	fmt.Println(list.Size, list)
//}
