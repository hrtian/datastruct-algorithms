package main

import (
	"bytes"
	"fmt"
)

const DefaultCapacity int = 8
const ElementNotFound int = -1


// 优化思路
/**
 * 减少 ArrayList的元素移动次数
 * 添加 first index指标
/
 */

/*
* size 元素个数
* capacity 数组容量
* element 元素
 */
type ArrayList struct {
	Size     int
	Elements []interface{}
}

// ArrayList construct function
func NewArrayList(capacity int) *ArrayList {
	if capacity <= DefaultCapacity {
		capacity = DefaultCapacity
	}

	return &ArrayList{
		Size:     0,
		Elements: make([]interface{}, capacity),
	}

}

func (a ArrayList) size() int {
	/*
	* return size of arraylist
	* @return
	 */
	return a.Size
}

func (a ArrayList) isEmpty() bool {
	return a.Size == 0
}

// O(1)
func (a ArrayList) get(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.Size))
	}
	return a.Elements[index]
}

// O(1)
func (a ArrayList) set(index int, element interface{}) interface{} {
	if index < 0 || index >= a.Size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.Size))
	}

	old := a.Elements[index]
	a.Elements[index] = element
	return old
}

// 待完善！！
func (a ArrayList) indexOf(element interface{}) int {

	for i := 0; i < a.Size; i++ {
		if a.Elements[i] == element {
			return i
		}
	}

	return ElementNotFound
}

func (a ArrayList) contains(element interface{}) bool {
	if a.indexOf(element) != ElementNotFound {
		return true
	}

	return false
}

func (a ArrayList) clear() {
	// clear all member in arraylist
	// 留下可循环利用的内存，置空不可循环利用的；
	for i := 0; i < a.Size; i++ {
		a.Elements[i] = nil
	}

	a.Size = 0
}

// O(1)  O(n)
func (a *ArrayList) removeIndex(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.Size))
	}

	src := a.Elements[index]
	for i := index + 1; i < a.Size; i++ {
		a.Elements[i-1] = a.Elements[i]
	}

	a.Elements[a.Size-1] = nil
	a.Size--
	a.trim()

	return src
}

func (a *ArrayList) removeElement(element interface{}) {
	a.removeIndex(a.indexOf(element))
}

// O(1)  O(n)
func (a *ArrayList) addIndex(index int, element interface{}) {
	if index < 0 || index > a.Size { // 判断不一样
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.Size))
	}

	a.ensureCapacity(a.Size + 1)

	for i := a.Size; i > index; i-- {
		a.Elements[i] = a.Elements[i-1]
	}
	a.Elements[index] = element
	a.Size++
}

// O(1) O(n) O(1) 均摊复杂度O(1)
// 经过连续多次复杂度比较低的情况后，出现个别复杂度搞得情况
func (a *ArrayList) add(element interface{}) {
	a.addIndex(a.Size, element)
}

func (a *ArrayList) ensureCapacity(capacity int) {
	oldCapacity := len(a.Elements)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := oldCapacity + oldCapacity>>1
	newElements := make([]interface{}, newCapacity)
	fmt.Println("扩容：", a.Size, oldCapacity, newCapacity)
	for i := 0; i < a.Size; i++ {
		newElements[i] = a.Elements[i]
	}

	a.Elements = newElements
}

// 内存使用紧张进行缩容，但是当扩容和缩容时机不恰当，会导致，时间复杂度震荡，（扩容倍数 * 缩容倍数 = 1）
func (a *ArrayList) trim() {

	oldCapacity := len(a.Elements)
	newCapacity := len(a.Elements)>>1
	if a.Size > oldCapacity >> 1  || oldCapacity <= DefaultCapacity {
		return
	}


	newElements := make([]interface{}, newCapacity)

	fmt.Println("缩容：", a.Size, len(a.Elements), newCapacity)
	for i := 0; i < a.Size; i++ {
		newElements[i] = a.Elements[i]
	}

	a.Elements = newElements
}

func (a *ArrayList) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("ArrayList: size = %d ", a.Size))
	buffer.WriteString("[")
	for i := 0; i < a.Size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(a.Elements[i]))
	}

	buffer.WriteString("]")

	return buffer.String()
}

func main() {
	list := NewArrayList(1)
	for i := 0; i < 50; i++ {
		list.add(i)
	}

	for i := 49; i >= 0; i-- {
		list.removeIndex(i)
	}
}
