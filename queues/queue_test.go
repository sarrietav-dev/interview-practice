package queues_test

import (
	"testing"

	"github.com/sarrietav-dev/interview-practice/queues"
)

func TestQueue(t *testing.T) {
    q := &queues.Queue[int]{}

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    if q.Dequeue() != 1 {
        t.Errorf("Expected 1, but got %v", q.Dequeue())
    }

    if q.Dequeue() != 2 {
        t.Errorf("Expected 2, but got %v", q.Dequeue())
    }

    if q.Dequeue() != 3 {
        t.Errorf("Expected 3, but got %v", q.Dequeue())
    }

	q.Enqueue(1)
}

func TestEmptyQueue(t *testing.T) {
    q := &queues.Queue[int]{}

    if !q.IsEmpty() {
        t.Errorf("Expected queue to be empty, but it's not")
    }
}

func TestPanicOnEmptyQueue(t *testing.T) {
    q := &queues.Queue[int]{}

    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected panic, but it didn't happen")
        }
    }()

    q.Dequeue()
}

func TestTailEmptyAfterDequeue(t *testing.T) {
    q := &queues.Queue[int]{}

    q.Enqueue(1)
	q.Dequeue()

    if !q.IsEmpty() {
        t.Errorf("Expected queue to be empty, but it's not")
    }
}