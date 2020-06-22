package bst

import (
	utils "code.golang.com/Datastruct/Utils"
	"errors"
)

// node is a tree node, export for travel.
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
		ret = b.compare(e, tmpNode.val)
		parentNode = tmpNode

		if ret > 0 {
			tmpNode = tmpNode.right
		} else if ret < 0 {
			tmpNode = tmpNode.left
		} else {
			// 当添加为自定义类型时，避免直接返回导致的添加元素的比较内容一致，而其他内容不一致
			tmpNode.val = e
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

// Remove is a function for remove element from BST
//func (b *BST) Remove(e interface{}) {
//	// todo
//}

//func (b *BST) Contains(e interface{}) bool {
//	// todo
//}

func (b *BST) compare(e1, e2 interface{}) int {
	if b.comparator == nil {
		return utils.Comparator(e1, e2)
	}

	return b.comparator(e1, e2)
}

// PreRecur is a function for travel tree by root-in-first
/*
 * @return void
 * @params visit func
 * @func recursion
 */
func (b *BST) PreRecur(visit func(interface{})) {

	var preRecur func(*node, func(interface{}))
	preRecur = func(root *node, visit func(interface{})) {
		if root == nil {
			return
		}

		visit(root.val)
		preRecur(root.left, visit)
		preRecur(root.right, visit)
	}

	preRecur(b.root, visit)
}

// InRecur is a function for travel element by root-in-middle
func (b *BST) InRecur(visit func(interface{})) {
	var inRecur func(*node, func(interface{}))
	inRecur = func(root *node, visit func(interface{})) {
		if root == nil {
			return
		}

		inRecur(root.left, visit)
		visit(root.val)
		inRecur(root.right, visit)
	}

	inRecur(b.root, visit)
}

// PostRecur is a function for travel element by root-in-last
func (b *BST) PostRecur(visit func(interface{})) {

	var postRecur func(*node, func(interface{}))
	postRecur = func(root *node, visit func(interface{})) {
		if root == nil {
			return
		}

		postRecur(root.left, visit)
		postRecur(root.right, visit)
		visit(root.val)
	}

	postRecur(b.root, visit)
}

// LevelTravel is a function for travel element by level.
func (b *BST) LevelTravel(visit func(interface{}) bool) {
	if b.root == nil || visit == nil {
		return
	}

	queue := []*node{b.root}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		if visit(v.val) {
			return
		}
		
		if v.left != nil {
			queue = append(queue, v.left)
		}
		if v.right != nil {
			queue = append(queue, v.right)
		}
	}
}
