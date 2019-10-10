package test_practice

import (
	"fmt"
	"testing"
)

// 判断链表是否有环
func TestLinkHasRing(t *testing.T) {
	var node1 = &Node{nil, 1}
	var node2 = &Node{nil, 2}
	var node3 = &Node{nil, 3}
	var node4 = &Node{nil, 4}
	var node5 = &Node{nil, 5}
	l := NewLinkList()
	l.Insert(node1).Insert(node2).Insert(node3).Insert(node4).Insert(node5)
	node5.Next = node2

	fmt.Println(l.HasRing())
}
