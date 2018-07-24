/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteFromList(lists []*ListNode, index int) []*ListNode {
	if index == len(lists)-1 {
		lists = lists[:index]
	} else {
		lists = append(lists[:index], lists[index+1:len(lists)]...)
	}
	return lists
}

func mergeKLists(lists []*ListNode) *ListNode {
	newLists := make([]*ListNode, 0, len(lists))
	for _, n := range lists {
		if n != nil {
			newLists = append(newLists, n)
		}
	}
	result := new(ListNode)
	pointer := result
	var index, minVal int

	for len(newLists) > 0 {
		minVal = newLists[0].Val
		index = 0
		for i, n := range newLists {
			if n != nil && n.Val < minVal {
				minVal = n.Val
				index = i
			}
		}
		pointer.Next = newLists[index]
		if newLists[index].Next == nil {
			newLists = deleteFromList(newLists, index)
		} else {
			newLists[index] = newLists[index].Next
		}
		pointer = pointer.Next
	}
	return result.Next
}

func nodesToList(lists []*ListNode) [][]int {
	result := make([][]int, 0)
	for _, l := range lists {
		temp := nodeToList(l)
		result = append(result, temp)
	}
	return result
}

func listsToNodeLists(lists [][]int) []*ListNode {
	result := make([]*ListNode, 0)
	for _, l := range lists {
		node := listToNode(l)
		result = append(result, node)
	}
	return result
}

func TestNodesToList(t *testing.T) {
	assert := assert.New(t)
	temp := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	temp2 := ListNode{2, &ListNode{5, nil}}
	assert.Equal(
		[][]int{[]int{1, 2, 3}, []int{2, 5}},
		nodesToList([]*ListNode{&temp, &temp2}),
	)
}

func TestListsToNodeLists(t *testing.T) {
	assert := assert.New(t)
	temp := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	temp2 := ListNode{2, &ListNode{5, nil}}
	assert.Equal(
		[]*ListNode{&temp, &temp2},
		listsToNodeLists([][]int{[]int{1, 2, 3}, []int{2, 5}}),
	)
}

func TestMergeKLists(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]int{1},
		nodeToList(mergeKLists(listsToNodeLists([][]int{[]int{}, []int{1}}))),
	)
}

// func mergeKLists(lists []*ListNode) *ListNode {
//     if lists == nil {
//         return nil
//     }

//     listLen := len(lists)

//     var heap = make([]*ListNode, 0, listLen)
//     for i := 0; i < listLen; i++ {
//         top := lists[i]
//         if top == nil {
//             continue
//         }
//         heap = append(heap, top)
//     }
//     for i := len(heap) - 1; i >= 0; i-- {
//         heap = adjust(heap, i)
//     }

//     var head *ListNode
//     var tail *ListNode

//     for len(heap) > 0 {
//         top := heap[0]
//         if head == nil {
//             head = top
//         } else {
//             tail.Next = top
//         }
//         tail = top
//         heap[0] = heap[0].Next
//         if heap[0] == nil {
//             heap[0] = heap[len(heap) - 1]
//             heap = heap[:len(heap)-1]
//         }
//         if len(heap) > 0 {
//             heap = adjust(heap, 0)
//         }
//     }
//     return head
// }

// func adjust(heap []*ListNode, i int) []*ListNode {
//     n := len(heap)
//     for i < n {
//         node := heap[i]
//         left := i * 2 + 1;
//         right := i * 2 + 2;
//         min := i
//         if left < n && (heap[left].Val < heap[min].Val) {
//             min = left
//         }
//         if right < n && (heap[right].Val < heap[min].Val) {
//             min = right
//         }
//         if min != i {
//             heap[i] = heap[min]
//             heap[min] = node
//             i = min
//         } else {
//             break
//         }
//     }
//     return heap
// }
