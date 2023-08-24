package queue

// Queue struct which represents the queue object
type Queue struct {
	Front     *Node
	Rear      *Node
	QueueSize int
}

// Node struct which represents an item in the queue
type Node struct {
	Value    int
	Previous *Node
}

// New create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Size returns the queue size (how many elements are currently in the queue)
func (q *Queue) Size() int {
	return q.QueueSize
}

// Enqueue add new node to the rear of the queue
func (q *Queue) Enqueue(data int) {
	n := &Node{data, nil}
	if q.Front == nil && q.Rear == nil {
		q.Front = n
		q.Rear = n
	} else {
		q.Rear.Previous = n
		q.Rear = n
	}

	q.QueueSize++
}

// Dequeue remove front node and return
func (q *Queue) Dequeue() (int, bool) {
	if q.Front == nil {
		return 0, false
	}

	n := q.Front
	if q.QueueSize == 1 {
		q.Front = nil
		q.Rear = nil
	} else {
		q.Front = n.Previous
	}

	q.QueueSize--
	return n.Value, true
}

// Peek peek front node value without removing from queue
func (q *Queue) Peek() (int, bool) {
	if q.Front == nil {
		return 0, false
	}
	return q.Front.Value, true
}
