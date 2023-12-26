package buffer

import (
	"fmt"
)

var (
	ErrBufferIsFull  = fmt.Errorf("buffer is full")
	ErrBufferIsEmpty = fmt.Errorf("buffer is empty")
)

type CircularBuffer struct {
	taskQueue [][]byte
	capacity  int
	head      int
	tail      int
	full      bool
}

func (s *CircularBuffer) IsEmpty() bool {
	return s.head == s.tail && !s.full
}

func (s *CircularBuffer) IsFull() bool {
	return s.full
}

func (s *CircularBuffer) Size() int {
	if s.full {
		return s.capacity
	}

	if s.tail >= s.head {
		return s.tail - s.head
	}

	return s.capacity - (s.head - s.tail)
}

func (s *CircularBuffer) Enqueue(task []byte) error {
	if s.IsFull() {
		return ErrBufferIsFull
	}

	s.taskQueue[s.tail] = task
	s.tail = (s.tail + 1) % s.capacity
	s.full = s.head == s.tail

	return nil
}

func (s *CircularBuffer) Dequeue() ([]byte, error) {
	if s.IsEmpty() {
		return nil, ErrBufferIsEmpty
	}

	data := s.taskQueue[s.head]
	s.full = false
	s.head = (s.head + 1) % s.capacity

	return data, nil
}

func (s *CircularBuffer) DequeueBatch(count int) ([][]byte, error) {
	if count < 0 {
		return nil, fmt.Errorf("invalid count")
	}

	if s.IsEmpty() {
		return nil, ErrBufferIsEmpty
	}

	if count > s.Size() {
		count = s.Size()
	}

	data := make([][]byte, count)
	for i := 0; i < count; i++ {
		data[i] = s.taskQueue[s.head]
		s.head = (s.head + 1) % s.capacity
	}

	s.full = false

	return data, nil
}

func NewCircularBuffer(size int) *CircularBuffer {
	w := &CircularBuffer{
		taskQueue: make([][]byte, size),
		capacity:  size,
	}

	return w
}
