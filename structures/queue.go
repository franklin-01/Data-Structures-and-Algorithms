package structures

type Queue struct {
	data []int
}

func (queue *Queue) Enqueue(value int) {
	queue.data = append(queue.data, value)
}

func (queue *Queue) Dequeue() int {
	length := len(queue.data)

	if length < 1 {
		panic("queue is empty")
	}

	first := queue.data[0]
	queue.data = queue.data[1:length]

	return first
}

func (queue *Queue) Peek() int {
	length := len(queue.data)

	if length < 1 {
		panic("queue is empty")
	}

	return queue.data[0]

}

func (queue *Queue) isEmpty() bool {
	length := len(queue.data)

	if length > 0 {
		return false
	}
	return true
}

func (queue *Queue) size() int {
	length := len(queue.data)

	return length

}
