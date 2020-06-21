package bst

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	utils "code.golang.com/Datastruct/Utils"
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

// BST is a data struct of tree, binary search tree
type BST struct {
	comparator func(interface{}, interface{}) int
	root       *node
	size       int
}

// New is a construct function for BST, you shoulld specify a comparator
func New(comparator func(interface{}, interface{}) int) *BST {
	return &BST{
		comparator: comparator,
	}
}

// Size is function for getting size of bst
func (b *BST) Size() int {
	return b.size
}

// IsEmpty is function for that it is empty of bst
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// Clear is a function for clear BST elements
func (b *BST) Clear() {
	b.size = 0
}

// Add is a function for add element to BST
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

	for tmpNode != nil {
		ret = b.comparate(e, tmpNode.val)
		parentNode = tmpNode

		if ret > 0 {
			tmpNode = tmpNode.right
		} else if ret < 0 {
			tmpNode = tmpNode.left
		} else {
			tmpNode.val = e
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

// Remove is a function for remove element from BST
func (b *BST) Remove(e interface{}) {
	// todo
}

//func (b *BST) Contains(e interface{}) bool {
//	// todo
//}

func (b *BST) comparate(e1, e2 interface{}) int {
	if b.comparator == nil {
		return utils.Comparator(e1, e2)
	}

	return b.comparator(e1, e2)
}