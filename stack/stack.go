package stack

type Stack struct {
	data []any
}

func (stack *Stack) Push(data any) {
	stack.data = append(stack.data, data)
}

func (stack *Stack) Pop() any {
	count := len(stack.data)
	if count == 0 {
		return nil
	}

	lastItem := stack.data[count-1]

	stack.data = stack.data[:count-1]
	return lastItem
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.data) == 0
}

func (stack *Stack) Reset() {
	stack.data = nil
}

func (stack *Stack) Dump() []any {
	var copiedStack = make([]any, len(stack.data))
	copy(copiedStack, stack.data)
	return copiedStack
}

func (stack *Stack) Peek() any {
	if len(stack.data) == 0 {
		return nil
	}

	return stack.data[len(stack.data)-1]
}
