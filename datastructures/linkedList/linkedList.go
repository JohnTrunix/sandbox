package linkedList

// LinkedList struct which represents the head
type LinkedList struct {
	HEAD *Node
}

// Node struct which represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// New create new linked list
func New() *LinkedList {
	return &LinkedList{nil}
}

// Add a new node to the linked list at the end
func (l *LinkedList) Add(data int) {
	node := &Node{Value: data, Next: nil}

	if l.HEAD == nil {
		l.HEAD = node
		return
	}

	curr := l.HEAD
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
}

// Search if node with value exists in linked list
func (l *LinkedList) Search(data int) bool {
	curr := l.HEAD

	if curr.Value == data {
		return true
	}

	for curr.Next != nil {
		curr = curr.Next

		if curr.Value == data {
			return true
		}
	}

	return false
}
