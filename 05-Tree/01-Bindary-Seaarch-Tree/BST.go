package bst

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	utils "code.golang.com/Datastruct/Utils"
)

// Visitor is a struct using as a visitor function
type Visitor struct {
	Visit func(e interface{}) bool
	Stop  bool
}

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

// New is a construct function for BST, you should specify a comparator
func New(comparator func(interface{}, interface{}) int) *BST {
	return &BST{
		comparator: comparator,
	}
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

func (b *BST) Remove(e interface{}) bool {
	return false
}

func (b *BST) Update(e interface{}) {

}

func (b *BST) Search(e interface{}) *node {
	return nil
}

// Size is function for getting size of bst
func (b *BST) Size() int {
	return b.size
}

// Clear is a function for clear BST elements
func (b *BST) Clear() {
	b.size = 0
}

// IsEmpty is function for that it is empty of bst
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// IsLeaf is a for judging wheather the node is a leaf-node
func IsLeaf(node *node) bool {
	return node.left == nil && node.right == nil
}

// Has2Children is for judging wheather the node has two child-node
func Has2Children(node *node) bool {
	return node.left != nil && node.right != nil
}

/*
 * travel elements in tree through four different functions
 * @ author: hirah
 * @ return: void
 * @ params: *Visitor
 * @ func: PreRecur InRecur PostRecur LevelTravel
 */

// PreRecur is a function for travel tree by root-in-first
func (b *BST) PreRecur(V *Visitor) {
	if V.Visit == nil || V.Stop || b.root == nil {
		return
	}

	var preRecur func(*node, *Visitor)
	preRecur = func(node *node, V *Visitor) {
		// stop recursion
		if node == nil || V.Stop {
			return
		}

		V.Stop = V.Visit(node.val)
		preRecur(node.left, V)
		preRecur(node.right, V)
	}

	preRecur(b.root, V)
}

// InRecur is a function for travel element by root-in-middle
func (b *BST) InRecur(V *Visitor) {
	if V.Visit == nil || V.Stop || b.root == nil {
		return
	}

	var inRecur func(*Visitor, *node)
	inRecur = func(V *Visitor, node *node) {
		// stop recursion
		if node == nil || V.Stop {
			return
		}

		inRecur(V, node.left)
		// stop processing
		if V.Stop {
			return
		}
		V.Stop = V.Visit(node.val)
		inRecur(V, node.right)
	}

	inRecur(V, b.root)
}

// PostRecur is a function for travel element by root-in-last
func (b *BST) PostRecur(V *Visitor) {
	if V.Visit == nil || V.Stop || b.root == nil {
		return
	}

	var postRecur func(*Visitor, *node)
	postRecur = func(V *Visitor, node *node) {
		// stop recursion
		if node == nil || V.Stop {
			return
		}

		postRecur(V, node.left)
		postRecur(V, node.right)

		// stop processing
		if V.Stop {
			return
		}
		V.Stop = V.Visit(node.val)
	}

	postRecur(V, b.root)
}

// LevelTravel is a function for travel element by level.
func (b *BST) LevelTravel(V *Visitor) {
	if V.Visit == nil || V.Stop || b.root == nil {
		return
	}

	queue := []*node{b.root}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		if V.Stop = V.Visit(v.val); V.Stop {
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

// GetHByRecur is a function for get depth of binary tree through Recursion.
func (b *BST) GetHByRecur() int {
	var height func(*node) int
	height = func(node *node) int {
		if node == nil {
			return 0
		}
		return 1 + int(math.Max(float64(height(node.left)), float64(height(node.right))))
	}

	return height(b.root)
}

// GetHByIota is a function for get depth of binary tree through Iota
func (b *BST) GetHByIota() int {
	if b.root == nil {
		return 0
	}

	tmp, height, level := []*node{b.root}, 0, 1

	for len(tmp) != 0 {
		// Queue processing
		v := tmp[0]
		tmp = tmp[1:]
		level--

		if v.left != nil {
			tmp = append(tmp, v.left)
		}

		if v.right != nil {
			tmp = append(tmp, v.right)
		}

		if level >= 0 {
			height++
			level = len(tmp)
		}
	}

	return height
}

// IsComplete is a function for judging wheather the binary tree is a complete-binary-tree
func (b *BST) IsComplete() bool {
	if b.root == nil {
		return false
	}

	/*
	* left != nil enQueue
	* left == nil, right != nil false
	* right != nil enQueue
	* left != nil, right == nil
	* 							 ------> right == nil, 判断是否为叶子节点
	* left == nil, right == nil
	 */
	tmp := []*node{b.root}
	for len(tmp) > 0 {
		v := tmp[0]
		tmp = tmp[1:]

		if v.left != nil {
			tmp = append(tmp, v.left)
		} else if v.right != nil {
			return false
		}

		if v.left != nil {
			tmp = append(tmp, v.right)
		} else {
			for _, v := range tmp {
				if !IsLeaf(v) {
					return false
				}
			}
		}
	}
	return true
}

// String is a function for printing Binary-Tree.
func (b *BST) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(fmt.Sprintf("BST:\n size: %d\n", b.size))

	var toString func(*node, *bytes.Buffer, string)
	toString = func(root *node, buffer *bytes.Buffer, prefix string) {
		if root == nil {
			return
		}

		toString(root.left, buffer, prefix+"---")
		buffer.WriteString(prefix + fmt.Sprintf("%v\n", root.val))
		// 注意递归时的buffer调用
		toString(root.right, buffer, prefix+"---")
	}

	toString(b.root, buffer, "")
	return buffer.String()
}

func (b *BST) compare(e1, e2 interface{}) int {
	if b.comparator == nil {
		return utils.Comparator(e1, e2)
	}

	return b.comparator(e1, e2)
}

// ReverseBT is a function for reverse the left-node and right-node of parent-node
func ReverseBT(b *BST) *BST {
	var reverse func(*node)
	reverse = func(root *node) {
		if root == nil {
			return
		}
		/*
			preorder:
			root.left, root.right = root.right, root.left
			reverse(root.left)
			reverse(root.right)

			inorder:
			reverse(root.left)
			root.left, root.right = root.right, root.left
			reverse(root.left)

			postorder:
			reverse(root.left)
			reverse(root.right)
			root.left, root.right = root.right, root.left
		*/

		tmp := []*node{root}

		for len(tmp) > 0 {
			v := tmp[0]
			tmp = tmp[1:]

			v.left, v.right = v.right, v.left
			if v.left != nil {
				tmp = append(tmp, v.left)
			}
			if v.right != nil {
				tmp = append(tmp, v.right)
			}
		}

	}

	reverse(b.root)
	return b
}
