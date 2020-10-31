package Queue

import (
	"bytes"
	"code.golang.com/Datastruct/02-LinkedList/03_DLL"
	"fmt"
)

type Queue struct {
	data *dll.DLL
}

func New() *Queue {
	return &Queue{
		data: dll.New(),
	}
}

func (q *Queue) Size() int {
	return q.data.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue) EnQueue(e interface{}) {
	q.data.Add(e)
}

func (q *Queue) DeQueue() interface{} {
	return q.data.Remove(0)
}

func (q *Queue) Front() interface{} {
	return q.data.Get(0)
}

func (q *Queue) Clear() {
	q.data.Clear()
}

func (q *Queue) String() string {
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
