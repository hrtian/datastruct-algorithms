package VHLL

import (
	"bytes"
	"fmt"
)

// 链表需要注重边界条件的测试 0 size size - 1, 解决 0 的方法为虚拟头节点

var ENotFound = -1

// LinkedList_vh is a signal linked VNode;
type LinkedListVh struct {
	Size  int
	first *VNode
}

type VNode struct {
	e    interface{}
	next *VNode
}

func New() *LinkedListVh {
	l := new(LinkedListVh)
	l.first = new(VNode)
	l.first.next = nil
	l.first.e = nil
	return l
}

func (l *LinkedListVh) clear() {
	l.Size = 0
	l.first = nil
}

func (l *LinkedListVh) size() int {
	return l.Size
}

func (l *LinkedListVh) isEmpty() bool {
	return l.Size == 0
}

func (l *LinkedListVh) contains(e interface{}) bool {
	return l.indexOf(e) != ENotFound
}

// O(1)  O(n)
func (l *LinkedListVh) get(index int) *VNode {
	return l.findNode(index)
}

// O(1) O(n)
func (l *LinkedListVh) set(index int, e interface{}) interface{} {
	node := l.findNode(index)
	old := node.e
	node.e = e
	return old
}

// O(1)
func (l *LinkedListVh) add(e interface{}) {
	l.addIndex(l.Size, e)
}

// override the addIndex function O(1)  O(n)
func (l *LinkedListVh) addIndex(index int, e interface{}) {
	var prev *VNode
	if index == 0 {
		prev = l.first
	} else {
		prev = l.findNode(index - 1)
	}

	next := prev.next
	prev.next = &VNode{
		e:    e,
		next: next,
	}

	l.Size++
}

// override the remove function O(1)  O(n)
func (l *LinkedListVh) remove(index int) interface{} {
	var prev *VNode

	if index == 0 {
		prev = l.first
	} else {
		prev = l.findNode(index - 1)
	}

	old := prev.next.e
	prev.next = prev.next.next

	l.Size--
	return old
}

// override the indexOf function
func (l *LinkedListVh) indexOf(e interface{}) int {
	node := l.first.next

	for i := 0; i < l.Size; i++ {
		if node.e == e {
			return i
		}
		node = node.next
	}
	return ENotFound
}

// override the findNode function
func (l *LinkedListVh) findNode(index int) *VNode {
	if index < 0 || index >= l.Size {
		panic(fmt.Sprintf("index = %d, size = %d, OutOfRange", index, l.Size))
	}

	node := l.first.next

	for ; index > 0; index-- {
		node = node.next
	}
	return node
}

// override the  String() function
func (l *LinkedListVh) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: LinkedList_vh, size: %d, [", l.Size))
	node := l.first.next

	for i := 0; i < l.Size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(node.e))
		node = node.next
	}
	buffer.WriteString("]")
	return buffer.String()
}

//func main() {
//	list := newLinkedList()
//	list.add(0)
//	list.add(1)
//	list.remove(0)
//	fmt.Println(list.indexOf(1))
//	list.add(2)
//	list.add(3)
//	fmt.Println(list)
//}
