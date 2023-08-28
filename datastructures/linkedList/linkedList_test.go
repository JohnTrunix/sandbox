package linkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	l := New()

	assert.Nil(t, l.HEAD)
}

func TestAdd(t *testing.T) {
	l := New()
	l.Add(10)
	l.Add(20)

	p := l.HEAD
	assert.Equal(t, 10, p.Value)

	p = p.Next
	assert.Equal(t, 20, p.Value)

	l.Add(30)
	p = p.Next
	assert.Equal(t, 30, p.Value)

	p = p.Next
	assert.Nil(t, p)
}
