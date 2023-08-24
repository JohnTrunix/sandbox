package binarySearchTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNode(t *testing.T) {
	root := New(100)

	assert.Equal(t, 100, root.Data)
	assert.Nil(t, root.LeftChild)
	assert.Nil(t, root.RightChild)
}

func TestInsertNode(t *testing.T) {
	root := New(100)
	root.Insert(50)
	root.Insert(150)

	assert.Equal(t, 100, root.Data)
	assert.Equal(t, 50, root.LeftChild.Data)
	assert.Equal(t, 150, root.RightChild.Data)
}

func TestInsertChildOfNode(t *testing.T) {
	root := New(100)
	root.Insert(50)
	root.Insert(150)
	root.Insert(40)
	root.Insert(160)
	root.Insert(55)
	root.Insert(145)

	assert.Equal(t, 100, root.Data)
	assert.Equal(t, 50, root.LeftChild.Data)
	assert.Equal(t, 150, root.RightChild.Data)
	assert.Equal(t, 40, root.LeftChild.LeftChild.Data)
	assert.Equal(t, 160, root.RightChild.RightChild.Data)
	assert.Equal(t, 55, root.LeftChild.RightChild.Data)
	assert.Equal(t, 145, root.RightChild.LeftChild.Data)
}

func TestSearchNode(t *testing.T) {
	root := New(100)
	root.Insert(50)
	root.Insert(150)
	root.Insert(40)
	root.Insert(160)
	root.Insert(55)
	root.Insert(53)
	root.Insert(145)

	s1, p1 := root.Search(100)
	assert.True(t, s1)
	assert.Equal(t, 50, p1.LeftChild.Data)
	assert.Equal(t, 150, p1.RightChild.Data)

	s2, p2 := root.Search(40)
	assert.True(t, s2)
	assert.Nil(t, p2.LeftChild)
	assert.Nil(t, p2.RightChild)

	s3, p3 := root.Search(55)
	assert.True(t, s3)
	assert.Equal(t, 53, p3.LeftChild.Data)
	assert.Nil(t, p3.RightChild)

	s4, p4 := root.Search(52)
	assert.False(t, s4)
	assert.Nil(t, p4)
}
