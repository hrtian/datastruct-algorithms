package Circle_DEQueue

import (
	"bytes"
	"fmt"
)

var DefaultCapacity int = 4

type CDEQueue struct {
	size  int
	front int
	data  []interface{}
}

func New(capacity int) *CDEQueue {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}
	return &CDEQueue{
		data: make([]interface{}, capacity),
	}
}

func (this *CDEQueue) Size() int {
	return this.size
}

func (this *CDEQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *CDEQueue) Front() interface{} {
	if this.size == 0 {
		panic("No element in CDEQueue")
	}
	return this.data[this.front]
}

func (this *CDEQueue) Rear() interface{} {
	if this.size == 0 {
		panic("No element in CDEQueue")
	}
	return this.data[this.index(this.size-1)]
}

func (this *CDEQueue) EnQueueFront(e interface{}) {
	// todo
	this.EnsureCapacity(this.size + 1)
	this.data[this.index(-1)] = e
	this.front = this.index(-1)
	this.size++

}

func (this *CDEQueue) EnQueueRear(e interface{}) {
	this.EnsureCapacity(this.size + 1)
	this.data[this.index(this.size)] = e
	this.size++
}

func (this *CDEQueue) DeQueueFront() interface{} {
	if this.size == 0 {
		panic("No element in CQueue")
	}
	old := this.data[this.front]
	this.data[this.front] = nil
	this.size--
	this.front = this.index(1)
	return old
}

func (this *CDEQueue) DeQueueRear() interface{} {
	if this.size == 0 {
		panic("No element in CDEQueue")
	}
	old := this.data[this.index(this.size-1)]
	this.data[this.index(this.size-1)] = nil
	this.size--
	return old
}

func (this *CDEQueue) EnsureCapacity(capacity int) {
	oldCapacity := len(this.data)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := oldCapacity + oldCapacity>>1
	newData := make([]interface{}, newCapacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[this.index(i)]
	}
	this.front = 0
	this.data = newData
}

func (this *CDEQueue) Clear() {
	this.size = 0
	for i := 0; i < len(this.data); i++ {
		this.data[i] = nil
	}
}

func (this *CDEQueue) index(idx int) int {
	if idx+this.front >= 0 {

		if idx+this.front >= len(this.data) {
			return idx + this.front - len(this.data)
		} else {
			return idx + this.front
		}

	} else {
		return idx + len(this.data)
	}
}

func (this *CDEQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: CQueue, size: %d, front: %d [", this.size, this.front))

	for i := 0; i < len(this.data); i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}

		buffer.WriteString(fmt.Sprint(this.data[i]))
	}

	buffer.WriteString("]")
	return buffer.String()
}
