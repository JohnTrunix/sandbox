package stack

// Stack struct which represents the stack object
type Stack struct {
	Head      *Node
	StackSize int
}

// Node struct which represents an item in the stack
type Node struct {
	Value    int
	Previous *Node
}

// New create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Size returns the stack size (how many elements are currently in the stack)
func (s *Stack) Size() int {
	return s.StackSize
}

// Push add a new node into the stack
func (s *Stack) Push(data int) {
	n := &Node{data, s.Head}
	s.Head = n
	s.StackSize++
}

// Pop remove the top node and return
func (s *Stack) Pop() (int, bool) {
	if s.Head == nil {
		return 0, false
	}

	p := s.Head
	s.Head = p.Previous
	s.StackSize--
	return p.Value, true
}

// Peek returns top node value without removing node from stack
func (s *Stack) Peek() (int, bool) {
	if s.Head == nil {
		return 0, false
	}

	return s.Head.Value, true
}
