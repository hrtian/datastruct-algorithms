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

// Node is a tree Node, export for travel.
type Node struct {
	val    interface{}
	left   *Node
	right  *Node
	parent *Node
}

func newNode(e interface{}, parent *Node) *Node {
	return &Node{
		val:    e,
		parent: parent,
	}
}

// BST is a data struct of tree, binary search tree
type BST struct {
	comparator func(interface{}, interface{}) int
	root       *Node
	size       int
}

// New is a construct function for BST, you should specify a comparator
func New(comparator func(interface{}, interface{}) int) *BST {
	return &BST{
		comparator: comparator,
	}
}

// AddByRecur is a function for add element to BST by recursion
func (b *BST) AddByRecur(e interface{}) error {
	if e == nil {
		return errors.New("element must not be nil")
	}

	var add func(*Node, interface{}) *Node
	add = func(node *Node, e interface{}) *Node {
		if node == nil {
			b.size++
			return &Node{
				val: e,
			}
		}
		cmp := b.compare(e, node.val)
		if cmp < 0 {
			node.left = add(node.left, e)
		} else if cmp > 0 {
			node.right = add(node.right, e)
		} else {
			node.val = e
		}
		b.size++
		return node
	}

	b.root = add(b.root, e)
	return nil
}

// AddByIota is a function for add element to BST by iota
func (b *BST) AddByIota(e interface{}) error {
	if e == nil {
		return errors.New("element must not be nil")
	}

	if b.root == nil {
		b.root = newNode(e, nil)
		b.size++
		return nil
	}

	parentNode, tmpNode, ret := new(Node), b.root, 0

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

// Remove is a function for delete element in BST
func (b *BST) Remove(e interface{}) bool {
	return false
}

// Update is a function for update node which val is old
func (b *BST) Update(old, new interface{}) {
	b.SearchByRec(old).val = new
}

// SearchByRec is a function for get node which val is 'val' by recursion.
func (b *BST) SearchByRec(e interface{}) *Node {
	var get func(*Node, interface{}) *Node
	get = func(node *Node, v interface{}) *Node {
		if node == nil {
			return nil
		}

		cmp := b.compare(e, node.val)
		if cmp < 0 {
			return get(node.left, e)
		} else if cmp > 0 {
			return get(node.right, e)
		} else {
			return node
		}
	}
	return get(b.root, e)
}

// SearchByIota is a function for get node which val is 'val' by recursion.
func (b *BST) SearchByIota(e interface{}) *Node {
	if b.root == nil {
		return nil
	}

	for tmpNode := b.root; tmpNode != nil; {
		cmp := b.compare(e, tmpNode.val)
		if cmp < 0 {
			tmpNode = tmpNode.left
		} else if cmp > 0 {
			tmpNode = tmpNode.right
		} else {
			return tmpNode
		}
	}
	return nil
}

// Contains is a function for judging wheather e is a val int BST
func (b *BST) Contains(e interface{}) bool {
	// return b.SearchByRec(e) != nil
	return b.SearchByIota(e) != nil
}

// Size is function for getting size of bst
func (b *BST) Size() int {
	return b.size
}

// GetSize is function for getting size of bst
func (b *BST) GetSize(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + b.GetSize(node.right) + b.GetSize(node.left)
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
func IsLeaf(node *Node) bool {
	return node.left == nil && node.right == nil
}

// Has2Children is for judging wheather the node has two child-node
func Has2Children(node *Node) bool {
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

	var preRecur func(*Node, *Visitor)
	preRecur = func(node *Node, V *Visitor) {
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

	var inRecur func(*Visitor, *Node)
	inRecur = func(V *Visitor, node *Node) {
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

	var postRecur func(*Visitor, *Node)
	postRecur = func(V *Visitor, node *Node) {
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

	queue := []*Node{b.root}
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

// String is a function for printing Binary-Tree.
func (b *BST) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(fmt.Sprintf("BST:\n size: %d\n", b.size))

	var toString func(*Node, *bytes.Buffer, string)
	toString = func(root *Node, buffer *bytes.Buffer, prefix string) {
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

// GetHByRecur is a function for get depth of binary tree through Recursion.
func (b *BST) GetHByRecur() int {
	var height func(*Node) int
	height = func(node *Node) int {
		if node == nil {
			return 0
		}
		return 1 + int(math.Max(float64(height(node.left)), float64(height(node.right))))
	}

	return height(b.root)
}

/********************************************************** 工具函数 *****************************************************************/

// GetHByIota is a function for get depth of binary tree through Iota
func (b *BST) GetHByIota() int {
	if b.root == nil {
		return 0
	}

	tmp, height, level := []*Node{b.root}, 0, 1
	for len(tmp) != 0 {
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
/*
* left != nil enQueue
* left == nil, right != nil false
* right != nil enQueue
* left != nil, right == nil
* 							 ------> right == nil, 判断是否为叶子节点
* left == nil, right == nil
 */
func (b *BST) IsComplete() bool {
	if b.root == nil {
		return false
	}

	tmp := []*Node{b.root}
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

// ReverseBT is a function for reverse the left-node and right-node of parent-node
func ReverseBT(b *BST) *BST {
	var reverse func(*Node)
	reverse = func(root *Node) {
		if root == nil {
			return
		}

		tmp := []*Node{root}

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

// IsValidBST is function for judging if a BT is a BST by recursion.
func (b *BST) IsValidBST() bool {
	var pre = math.MinInt64
	var isValid func(*Node) bool
	isValid = func(node *Node) bool {
		if node == nil {
			return true
		}
		if !isValid(node.left) {
			return false
		}
		if b.compare(pre, node.val) >= 0 {
			return false
		}
		return isValid(node.right)
	}

	return isValid(b.root)
}

// MinByRecur is a function for get min val of BST by recursion
func (b *BST) MinByRecur() interface{} {
	var min func(*Node) *Node
	min = func(node *Node) *Node {
		if node.left == nil {
			return node
		}

		return min(node.left)
	}
	return min(b.root).val
}

// MinByIota is a function for get min val of BST by iota
func (b *BST) MinByIota() interface{} {
	if b.root == nil {
		return nil
	}
	node := b.root
	for node.left != nil {
		node = node.left
	}
	return node.val
}

// MaxByRecur is a function for get max val of BST by recursion
func (b *BST) MaxByRecur() interface{} {
	var max func(*Node) *Node
	max = func(node *Node) *Node {
		if node.right == nil {
			return node
		}
		return max(node.right)
	}
	return max(b.root).val
}

// MaxByIota is a function for get max val of BST by iota
func (b *BST) MaxByIota() interface{} {
	if b.root == nil {
		return nil
	}
	node := b.root
	for node.right != nil {
		node = node.right
	}
	return node.val
}

// Floor <=
func (b *BST) Floor(e interface{}) *Node {
	var floor func(*Node, interface{}) *Node
	floor = func(node *Node, e interface{}) *Node {
		if node == nil {
			return nil
		}
		cmp := b.compare(e, node.val)
		if cmp == 0 {
			return node
		}
		if cmp > 0 {
			return floor(node.left, e)
		}

		ret := floor(node.right, e)
		if ret != nil {
			return ret
		}
		return node
	}

	node := floor(b.root, e)
	if node == nil {
		return nil
	}
	return node
}

// Ceil >=
func (b *BST) Ceil(e interface{}) *Node {
	var ceil func(*Node, interface{}) *Node
	ceil = func(node *Node, e interface{}) *Node {
		if node == nil {
			return nil
		}

		cmp := b.compare(e, node.val)
		if cmp == 0 {
			return node
		}
		if cmp > 0 {
			return ceil(node.right, e)
		}

		ret := ceil(node.left, e)
		if ret != nil {
			return ret
		}
		return node
	}

	node := ceil(b.root, e)
	if node == nil {
		return nil
	}
	return node
}

// Select is a function for find node which is more than K nodes
func (b *BST) Select(k int) interface{} {
	var selectK func(*Node, int) *Node
	selectK = func(node *Node, k int) *Node {
		if node == nil {
			return nil
		}

		t := b.GetSize(node.left)
		fmt.Println(t)
		if t > k {
			return selectK(node.left, k)
		} else if t < k {
			// k - t - 1 --> k - t 为向右需要的个数， - 1的原因是node.right 是一个大于目标的节点
			return selectK(node.right, k-t-1)
		} else {
			return node
		}

	}

	node := selectK(b.root, k)
	if node != nil {
		return	node.val
	}
	return nil
}


func(b *BST)Rank(e interface{}) int {
	var rank func(*Node, interface{}) int
	rank = func(node *Node, e interface{}) int {
		if node == nil {return 0}
		cmp := b.compare(e, node.val)

		if cmp < 0 {
			return rank(node.left, e)
		} else if cmp > 0 {
			return 1 + b.GetSize(node.left) + rank(node.right, e)
		} else {
			return b.GetSize(node.left)
		}
	}
	
	return rank(b.root, e)
}