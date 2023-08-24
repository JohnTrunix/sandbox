package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStack(t *testing.T) {
	s := New()

	assert.Nil(t, s.Head)
	assert.Equal(t, 0, s.StackSize)
}

func TestSize(t *testing.T) {
	s := New()
	assert.Equal(t, 0, s.Size())

	s.Push(10)
	s.Push(10)
	s.Push(20)
	assert.Equal(t, 3, s.Size())
}

func TestPush(t *testing.T) {
	s := New()
	s.Push(100)

	assert.Equal(t, 100, s.Head.Value)
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(100)
	s.Push(200)

	pop1, ok1 := s.Pop()
	assert.Equal(t, 200, pop1)
	assert.True(t, ok1)

	pop2, ok2 := s.Pop()
	assert.Equal(t, 100, pop2)
	assert.True(t, ok2)

	pop3, ok3 := s.Pop()
	assert.Equal(t, 0, pop3)
	assert.False(t, ok3)
}

func TestPeek(t *testing.T) {
	s := New()

	peek1, ok1 := s.Peek()
	assert.Equal(t, 0, peek1)
	assert.False(t, ok1)

	s.Push(100)
	peek2, ok2 := s.Peek()
	assert.Equal(t, 100, peek2)
	assert.True(t, ok2)

}
