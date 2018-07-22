/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	pointer := result
	l, r := l1, l2
	for l != nil && r != nil {
		if l.Val <= r.Val {
			pointer.Next = l
			l = l.Next
		} else {
			pointer.Next = r
			r = r.Next
		}
		pointer = pointer.Next
	}
	for l != nil {
		pointer.Next = l
		l = l.Next
		pointer = pointer.Next
	}
	for r != nil {
		pointer.Next = r
		r = r.Next
		pointer = pointer.Next
	}
	return result.Next
}

func nodeToList(node *ListNode) []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func listToNode(nums []int) *ListNode {
	result := new(ListNode)
	pointer := result
	for _, num := range nums {
		temp := &ListNode{num, nil}
		pointer.Next = temp
		pointer = pointer.Next
	}
	return result.Next
}

func TestNodeToList(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{1, 2, 3}, nodeToList(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}

func TestListToNode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, listToNode([]int{1, 2, 3}))
}
