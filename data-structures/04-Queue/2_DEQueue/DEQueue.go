package DEQueue

import (
	"bytes"
	"code.golang.com/Datastruct/02-LinkedList/03_DLL"
	"fmt"
)

type DEQueue struct {
	data *dll.DLL
}

func New() *DEQueue {
	return &DEQueue{
		data: dll.New(),
	}
}

func (q *DEQueue) Size() int {
	return q.data.Size()
}

func (q *DEQueue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *DEQueue) EnQueueRear(e interface{}) {
	q.data.Add(e)
}

func (q *DEQueue) EnQueueFront(e interface{}) {
	q.data.AddIndex(0, e)
}

func (q *DEQueue) DeQueueRear() interface{} {
	return q.data.Remove(q.data.Size() - 1)
}

func (q *DEQueue) DeQueueFront() interface{} {
	return q.data.Remove(0)
}

func (q *DEQueue) Front() interface{} {
	return q.data.Get(0)
}

func (q *DEQueue) rear() interface{} {
	return q.data.Get(q.data.Size() - 1)
}

func (q *DEQueue) Clear() {
	q.data.Clear()
}

func (q *DEQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type: Queue, size: %d, [", q.data.Size()))

	for i := 0; i < q.data.Size(); i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(q.data.Get(i)))
	}
	buffer.WriteString("]")
	return buffer.String()
}
