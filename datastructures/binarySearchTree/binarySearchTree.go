package binarySearchTree

import (
	"fmt"
)

// Node struct to represent a node in the binary search tree
type Node struct {
	Data       int
	LeftChild  *Node
	RightChild *Node
}

// New create a new node from struct for the binary search tree
func New(data int) *Node {
	return &Node{Data: data}
}

// Insert add new data as new node in binary search tree
func (n *Node) Insert(data int) {
	if data < n.Data {
		if n.LeftChild == nil {
			n.LeftChild = New(data)
		} else {
			n.LeftChild.Insert(data)
		}
	} else if data > n.Data {
		if n.RightChild == nil {
			n.RightChild = New(data)
		} else {
			n.RightChild.Insert(data)
		}
	} else {
		fmt.Printf("%d already in binary search tree\n", data)
	}
}

// Search try to find a value in binary search tree and returns success and pointer
func (n *Node) Search(data int) (bool, *Node) {
	if n == nil {
		return false, nil
	} else if data < n.Data {
		return n.LeftChild.Search(data)
	} else if data > n.Data {
		return n.RightChild.Search(data)
	}
	return true, n
}

// Delete try to find and delete a value
// Output create a printable json object
