package Circle_Queue

import (
	"bytes"
	"fmt"
)

const DefaultCapacity int = 4

type CQueue struct {
	front int
	size  int
	data  []interface{}
}

func New(capacity int) *CQueue {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}

	return &CQueue{
		data: make([]interface{}, capacity),
	}
}

func (cq *CQueue) Size() int {
	return cq.size
}

func (cq *CQueue) IsEmpty() bool {
	return cq.size == 0
}

func (cq *CQueue) Front() interface{} {
	if cq.size == 0 {
		return nil
	}
	return cq.data[cq.front]
}

func (cq *CQueue) EnQueue(e interface{}) {
	cq.EnsureCapacity(cq.size + 1)
	cq.data[cq.index(cq.size)] = e
	cq.size++
}

func (cq *CQueue) DeQueue() interface{} {
	if cq.size == 0 {
		panic(fmt.Sprintf("No element in CQueue"))
	}
	old := cq.data[cq.front]
	cq.data[cq.front] = nil
	cq.size--
	cq.front = cq.index(1)
	return old
}

func (cq *CQueue) EnsureCapacity(capacity int) {
	oldCapacity := len(cq.data)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := oldCapacity + oldCapacity>>1
	newData := make([]interface{}, newCapacity)
	for i := 0; i < cq.size; i++ {
		newData[i] = cq.data[cq.index(i)]
	}
	cq.front = 0
	cq.data = newData
}

func (cq *CQueue) index(idx int) int {
	if idx+cq.front >= len(cq.data) {
		return idx + cq.front - len(cq.data)
	} else {
		return idx + cq.front
	}
}

func (cq *CQueue) Clear() {
	cq.size = 0
	for i := 0; i < len(cq.data); i++ {
		cq.data[i] = nil
	}
}

func (cq *CQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: CQueue, size: %d, [", cq.size))

	for i := 0; i < len(cq.data); i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}

		buffer.WriteString(fmt.Sprint(cq.data[i]))
	}

	buffer.WriteString("]")
	return buffer.String()
}
