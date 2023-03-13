package algo

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	items := map[string]int{"a": 2, "b": 3, "c": 1, "d": 0, "f": 9, "s": 6}
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for k, v := range items {
		pq[i] = &Item[string]{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	for i := 0; i < len(pq)/2; i++ {
		left := i*2 + 1
		right := i*2 + 2
		if left < len(pq) {
			assert.True(t, pq[i].priority <= pq[left].priority, "minHeap[%d] > minHeap[%d]", i, left)
		}
		if right < len(pq) {
			assert.True(t, pq[i].priority <= pq[right].priority, "minHeap[%d] > minHeap[%d]", i, right)
		}
	}
}

func TestPush(t *testing.T) {
	items := map[string]int{"a": 2, "b": 3, "c": 1, "d": 0, "f": 9, "s": 6}
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for k, v := range items {
		pq[i] = &Item[string]{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	PushToPQ(&pq, &Item[string]{value: "g", priority: 4, index: len(pq)})

	for i := 0; i < len(pq)/2; i++ {
		left := i*2 + 1
		right := i*2 + 2
		if left < len(pq) {
			assert.True(t, pq[i].priority <= pq[left].priority, "minHeap[%d] > minHeap[%d]", i, left)
		}
		if right < len(pq) {
			assert.True(t, pq[i].priority <= pq[right].priority, "minHeap[%d] > minHeap[%d]", i, right)
		}
	}
}

func TestUpdate(t *testing.T) {
	items := map[string]int{"a": 2, "b": 3, "c": 1, "d": 0, "f": 9, "s": 6}
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for k, v := range items {
		pq[i] = &Item[string]{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	PushToPQ(&pq, &Item[string]{value: "g", priority: 4, index: len(pq)})
	pq.Update(&Item[string]{value: "g", priority: 4, index: len(pq) - 1}, 0)

	for i := 0; i < len(pq)/2; i++ {
		left := i*2 + 1
		right := i*2 + 2
		if left < len(pq) {
			assert.True(t, pq[i].priority <= pq[left].priority, "minHeap[%d] > minHeap[%d]", i, left)
		}
		if right < len(pq) {
			assert.True(t, pq[i].priority <= pq[right].priority, "minHeap[%d] > minHeap[%d]", i, right)
		}
	}
}
