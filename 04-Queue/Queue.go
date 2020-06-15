package Queue

import (
	"code.golang.com/Datastruct/02-LinkedList/03_DLL"
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

