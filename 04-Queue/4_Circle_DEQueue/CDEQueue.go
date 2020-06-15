package Circle_DEQueue

var DefaultCapacity int = 4
type CDEQueue struct {
	size int
	front int
	rear int
	data []interface{}
}

func New(capacity int) *CDEQueue {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}
	return &CDEQueue{
		data: make([]interface{}, capacity),
	}
}

func (cdq * CDEQueue) Size() int {
	return cdq.size
}

func (cdq * CDEQueue) IsEmpty() bool{
	return cdq.size == 0
}

func (cdq * CDEQueue) enQueueFront(e interface{}) {

}

func (cdq * CDEQueue) enQueueRear(e interface{}) {

}

func (cdq * CDEQueue) deQueueFront() interface{} {
	return nil
}

func (cdq * CDEQueue) deQueueRear() interface{} {
	return nil
}

func (cdq * CDEQueue) index(idx int) int {
	return -1
}