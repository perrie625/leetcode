package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	listLen := len(lists)

	var heap = make([]*ListNode, 0, listLen)
	for i := 0; i < listLen; i++ {
		top := lists[i]
		if top == nil {
			continue
		}
		heap = append(heap, top)
	}
	for i := len(heap) - 1; i >= 0; i-- {
		heap = adjust(heap, i)
	}

	var head *ListNode
	var tail *ListNode

	for len(heap) > 0 {
		top := heap[0]
		if head == nil {
			head = top
		} else {
			tail.Next = top
		}
		tail = top
		heap[0] = heap[0].Next
		if heap[0] == nil {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		}
		if len(heap) > 0 {
			heap = adjust(heap, 0)
		}
	}
	return head
}

func adjust(heap []*ListNode, i int) []*ListNode {
	n := len(heap)
	for i < n {
		node := heap[i]
		left := i*2 + 1
		right := i*2 + 2
		min := i
		if left < n && (heap[left].Val < heap[min].Val) {
			min = left
		}
		if right < n && (heap[right].Val < heap[min].Val) {
			min = right
		}
		if min != i {
			heap[i] = heap[min]
			heap[min] = node
			i = min
		} else {
			break
		}
	}
	return heap
}

func TestMergeKLists2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]int{1},
		nodeToList(mergeKLists2(listsToNodeLists([][]int{[]int{}, []int{1}}))),
	)
}
