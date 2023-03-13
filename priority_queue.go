package algo

import "container/heap"

// Create priority queue
type Item[T comparable] struct {
	value    T
	priority int
	index    int
}

// Min heap
type PriorityQueue[T comparable] []*Item[T]

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	item := old[len(old)-1]
	old[len(old)-1] = nil
	item.index = -1
	*pq = old[:len(old)-1]

	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue[T]) Push(item any) {
	i := item.(*Item[T])
	i.index = len(*pq)
	*pq = append(*pq, i)
}

func PushToPQ[T comparable](pq *PriorityQueue[T], item *Item[T]) {
	heap.Push(pq, item)
}
