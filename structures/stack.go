package structures

type Stack struct {
	data []int
}

func (stack *Stack) Push(value int) {
	stack.data = append(stack.data, value)
}

func (stack *Stack) Pop(value int) int {
	length := len(stack.data)

	if length < 1 {
		panic("stack is empty")
	}

	last := stack.data[length-1]
	stack.data = stack.data[:length-1]

	return last

}

func (stack *Stack) Peek() int {

	length := len(stack.data)

	if length < 1 {
		panic("stack is empty")
	}

	return stack.data[length-1]

}

func (stack *Stack) isEmpty() bool {
	length := len(stack.data)

	if length > 0 {
		return false
	}
	return true
}

func (stack *Stack) size() int {
	length := len(stack.data)

	return length

}
