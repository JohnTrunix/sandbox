# Linked List

A `Linked List` is a node-based structure which is linear. Each element (node) has a data and a next field which points to the next node (or nil at the end). 
Therefore, the nodes do not have to be stored at contiguous memory locations.

```mermaid
flowchart LR
    H[Head]
    Ni[Null]
    
    subgraph Node 1 
        D1[Data]
        N1[Next Pointer]
    end
    subgraph Node 2
        D2[Data]
        N2[Next Pointer]
    end
    subgraph Node 3
        D3[Data]
        N3[Next Pointer]
    end
    
    H --> D1
    N1 --> D2
    N2 --> D3
    N3 --> Ni
    
```

**Pointers:**

- `Head`: this is the entry point and points to first node
- `Next`: points to the next node

## Insert new Node in Linked List

_Iterates through list until `next` points to nil. Insert new node at the end of linked list._

```go
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
```

## Search Node in Linked List

_Iterates through list until `next` points to correct value or nil._

```go
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
```

## Remove Node form Linked List
