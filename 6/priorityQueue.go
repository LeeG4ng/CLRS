package main

import (
	"math"
)

type Element interface {
	key() float64
	setKey(float64)
}

type MinPriorityQueue struct {
	heap []Element // 不使用heap[0]，从heap[1]开始
	size int
}

func NewQueue() *MinPriorityQueue {
	return &MinPriorityQueue{
		heap: []Element{nil}, // 不使用第一个元素
		size: 0,
	}
}

func (q *MinPriorityQueue) ExtractMin() Element {
	if q.size < 1 {
		panic("heap underflow")
	}
	min := q.heap[1]
	q.heap[1] = q.heap[q.size] // 使用最后一个覆盖
	q.heap = q.heap[:q.size]   // 并删除最后一个
	q.size--
	q.minHeapify(1)
	return min
}

func (q *MinPriorityQueue) DecreaseKey(i int, key float64) {
	if key > q.heap[i].key() {
		panic("new key is larger than current key")
	}
	q.heap[i].setKey(key)
	for i > 1 && q.heap[parent(i)].key() > q.heap[i].key() {
		q.heap[i], q.heap[parent(i)] = q.heap[parent(i)], q.heap[i]
		i = parent(i)
	}
}

func (q *MinPriorityQueue) Insert(e Element) {
	key := e.key()
	e.setKey(math.MaxFloat64)
	q.size++
	q.heap = append(q.heap, e)
	q.DecreaseKey(q.size, key)
}

func (q *MinPriorityQueue) minHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l <= q.size && q.heap[l].key() < q.heap[smallest].key() {
		smallest = l
	}
	if r <= q.size && q.heap[r].key() < q.heap[smallest].key() {
		smallest = r
	}
	if smallest != i {
		q.heap[i], q.heap[smallest] = q.heap[smallest], q.heap[i]
		q.minHeapify(smallest)
	}
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}
