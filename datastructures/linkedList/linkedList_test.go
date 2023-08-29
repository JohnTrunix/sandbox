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

func TestSearch(t *testing.T) {
	l := New()
	l.Add(100)
	l.Add(200)
	l.Add(300)
	l.Add(400)

	r1 := l.Search(100)
	assert.True(t, r1)

	r2 := l.Search(200)
	assert.True(t, r2)

	r3 := l.Search(300)
	assert.True(t, r3)

	r4 := l.Search(400)
	assert.True(t, r4)

	r5 := l.Search(800)
	assert.False(t, r5)

}
