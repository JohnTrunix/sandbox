package main

import (
	"datastructures/binarySearchTree"
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
}
