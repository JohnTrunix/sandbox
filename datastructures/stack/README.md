# Stack

A `Stack` is a linear data structure with a fixed order. Last node pushed to stack is first node popped from stack.

```mermaid
graph TD
    A1[Stack]
    B1[Element 1]
    H1[HEAD]
    
    A1 --> B1
    H1 --> B1
    
    A2[Stack]
    B2[Element 1]
    C2[Element 2]
    H2[HEAD]
    
    A2 --> B2
    B2 --> C2
    H2 --> C2
```

**Pointers:**
- `Head`: points to last added node
- `Previous`: node pointer which points to 2nd last added node

## Push new data to stack

_Push data to stack and move `HEAD` pointer to added Node_

```go
func (s *Stack) Push(data int) {
	n := &Node{data, s.Head}
	s.Head = n
	s.StackSize++
}
```

## Pop data from stack (pop last added node)

_Pop data from stack and move `HEAD` pointer to second last added Node_

```go
func (s *Stack) Pop() (int, bool) {
	if s.Head == nil {
		return 0, false
	}

	p := s.Head
	s.Head = p.Previous
	s.StackSize--
	return p.Value, true
}
```

## Peek data from stack (peek last added node)

_Peek data from stack, get value from `HEAD` pointer_

```go
func (s *Stack) Peek() (int, bool) {
	if s.Head == nil {
		return 0, false
	}

	return s.Head.Value, true
}
```