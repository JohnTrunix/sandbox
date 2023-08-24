package main

import (
	"datastructures/binarySearchTree"
	"datastructures/stack"
	"fmt"
)

func main() {
	/*
		Binary Search Tree
	*/
	root := binarySearchTree.New(10)

	// create Nodes in Tree
	root.Insert(5)
	root.Insert(20)
	root.Insert(15)

	// search Node
	fmt.Println(root.Search(15))

	/*
		Stack
	*/
	s := stack.New()
	s.Push(100)
	s.Push(200)

	fmt.Println(s.Size())
	fmt.Println(s.Pop())
}
