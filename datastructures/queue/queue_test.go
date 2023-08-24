package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStack(t *testing.T) {
	q := New()

	assert.Nil(t, q.Front)
	assert.Nil(t, q.Rear)
	assert.Equal(t, 0, q.QueueSize)
}

func TestEnqueue(t *testing.T) {
	q := New()

	q.Enqueue(100)
	assert.Equal(t, 100, q.Front.Value)
	assert.Equal(t, 100, q.Rear.Value)

	q.Enqueue(200)
	assert.Equal(t, 100, q.Front.Value)
	assert.Equal(t, 200, q.Rear.Value)
}

func TestDequeue(t *testing.T) {
	q := New()
	q.Enqueue(100)
	q.Enqueue(200)

	dequeue1, ok1 := q.Dequeue()
	assert.Equal(t, 100, dequeue1)
	assert.True(t, ok1)

	dequeue2, ok2 := q.Dequeue()
	assert.Equal(t, 200, dequeue2)
	assert.True(t, ok2)

	dequeue3, ok3 := q.Dequeue()
	assert.Equal(t, 0, dequeue3)
	assert.False(t, ok3)
}

func TestPeek(t *testing.T) {
	q := New()

	peek1, ok1 := q.Peek()
	assert.Equal(t, 0, peek1)
	assert.False(t, ok1)

	q.Enqueue(100)
	peek2, ok2 := q.Peek()
	assert.Equal(t, 100, peek2)
	assert.True(t, ok2)
	assert.Equal(t, 1, q.QueueSize)
}
