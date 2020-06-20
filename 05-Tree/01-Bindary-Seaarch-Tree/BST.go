package BST

import (
	"code.golang.com/Datastruct/Utils"
	"errors"
)

type node struct {
	val    interface{}
	left   *node
	right  *node
	parent *node
}

func newNode(e interface{}, parent *node) *node {
	return &node{
		val:    e,
		parent: parent,
	}
}

type BST struct {
	size int
	root *node
}

func New() *BST {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Clear() {
	b.size = 0
}

func (b *BST) Add(e interface{}) error {
	if e == nil {
		return errors.New("element must not be nil")
	}

	if b.root == nil {
		b.root = newNode(e, nil)
		b.size++
		return nil
	}

	parentNode, tmpNode, ret := new(node), b.root, 0

	for ; tmpNode != nil; {
		ret = Utils.Compare(e, tmpNode.val)
		parentNode = tmpNode

		if ret > 0 {
			tmpNode = tmpNode.right
		} else if ret < 0 {
			tmpNode = tmpNode.left
		} else {
			return nil
		}
	}

	if ret == 1 {
		parentNode.right = newNode(e, parentNode)
	} else {
		parentNode.left = newNode(e, parentNode)
	}
	b.size++
	return nil

}

func (b *BST) Remove(e interface{}) {
	// todo
}

//func (b *BST) Contains(e interface{}) bool {
//	// todo
//}
