package arrayList

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
	size     int
	elements []interface{}
}

// ArrayList construct function
func New(capacity int) *ArrayList {
	if capacity <= DefaultCapacity {
		capacity = DefaultCapacity
	}

	return &ArrayList{
		elements: make([]interface{}, capacity),
	}

}

/*
* return size of arraylist
* @return
 */
func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}

// O(1)
func (a *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.size))
	}
	return a.elements[index]
}

// O(1)
func (a *ArrayList) Set(index int, element interface{}) interface{} {
	if index < 0 || index >= a.size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.size))
	}

	old := a.elements[index]
	a.elements[index] = element
	return old
}

// 待完善！！
func (a *ArrayList) IndexOf(element interface{}) int {

	for i := 0; i < a.size; i++ {
		if a.elements[i] == element {
			return i
		}
	}

	return ElementNotFound
}

func (a *ArrayList) Contains(element interface{}) bool {
	if a.IndexOf(element) != ElementNotFound {
		return true
	}

	return false
}

func (a *ArrayList) Clear() {
	// clear all member in arraylist
	// 留下可循环利用的内存，置空不可循环利用的；
	for i := 0; i < a.size; i++ {
		a.elements[i] = nil
	}

	a.size = 0
}

// O(1)  O(n)
func (a *ArrayList) RemoveIndex(index int) interface{} {
	if index < 0 || index >= a.size {
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.size))
	}

	src := a.elements[index]
	for i := index + 1; i < a.size; i++ {
		a.elements[i-1] = a.elements[i]
	}

	a.elements[a.size-1] = nil
	a.size--
	a.trim()

	return src
}

func (a *ArrayList) removeElement(element interface{}) {
	a.RemoveIndex(a.IndexOf(element))
}

// O(1)  O(n)
func (a *ArrayList) addIndex(index int, element interface{}) {
	if index < 0 || index > a.size { // 判断不一样
		panic(fmt.Sprintf("index: %d, size: %d, OutOfRange", index, a.size))
	}

	a.ensureCapacity(a.size + 1)

	for i := a.size; i > index; i-- {
		a.elements[i] = a.elements[i-1]
	}
	a.elements[index] = element
	a.size++
}

// O(1) O(n) O(1) 均摊复杂度O(1)
// 经过连续多次复杂度比较低的情况后，出现个别复杂度搞得情况
func (a *ArrayList) Add(element interface{}) {
	a.addIndex(a.size, element)
}

func (a *ArrayList) ensureCapacity(capacity int) {
	oldCapacity := len(a.elements)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := oldCapacity + oldCapacity>>1
	newElements := make([]interface{}, newCapacity)
	fmt.Println("扩容：", a.size, oldCapacity, newCapacity)
	for i := 0; i < a.size; i++ {
		newElements[i] = a.elements[i]
	}

	a.elements = newElements
}

// 内存使用紧张进行缩容，但是当扩容和缩容时机不恰当，会导致，时间复杂度震荡，（扩容倍数 * 缩容倍数 = 1）
func (a *ArrayList) trim() {

	oldCapacity := len(a.elements)
	newCapacity := len(a.elements)>>1
	if a.size > oldCapacity >> 1  || oldCapacity <= DefaultCapacity {
		return
	}


	newElements := make([]interface{}, newCapacity)

	fmt.Println("缩容：", a.size, len(a.elements), newCapacity)
	for i := 0; i < a.size; i++ {
		newElements[i] = a.elements[i]
	}

	a.elements = newElements
}

func (a *ArrayList) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("ArrayList: size = %d ", a.size))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(a.elements[i]))
	}

	buffer.WriteString("]")

	return buffer.String()
}

//func main() {
//	list := New(1)
//	for i := 0; i < 50; i++ {
//		list.add(i)
//	}
//
//	for i := 49; i >= 0; i-- {
//		list.removeIndex(i)
//	}
//}
