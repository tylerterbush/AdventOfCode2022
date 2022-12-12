package common

import "fmt"

// FIFO is a FIFO queue
type FIFO struct {
	queue []interface{}
}

// New creates new FIFO and returns it
func NewQueue() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

// Push pushed node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

func (f *FIFO) Empty() bool {
	return len(f.queue) == 0
}

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := NewQueue()

	for _, val := range vals {
		fifo.Push(val)
	}

	for {
		front := fifo.Front()
		if front == nil {
			break
		}
		fmt.Println(front)
	}
}