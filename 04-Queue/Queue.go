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

func (q *Queue) enQueue(e interface{}) {
	q.data.Add(e)
}

func (q *Queue) deQueue() interface{} {
	return q.data.Remove(0)
}

func (q *Queue) front() interface{} {
	q.data.Get(0)
}

func (q *Queue) clear() {
	q.data.Clear()
}

