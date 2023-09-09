package queues

type Queue[T comparable] struct {
	head *node[T]
	tail *node[T]
}

type node[T comparable] struct {
	value    T
	previous *node[T]
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &node[T]{value: value, previous: nil}

	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.previous = newNode

	q.tail = newNode
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("the queue is empty")
	}

	value := q.head.value
	q.head = q.head.previous

	if q.IsEmpty() {
		q.tail = nil
	}

	return value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}