package test_practice

import (
	"fmt"
	"testing"
)

// 链表的长度，不包过头
type Node struct {
	Next *Node
	Data int
}
type LinkList struct {
	Header *Node
}

func NewLinkList() *LinkList {
	return &LinkList{
		Header: &Node{
			Next: nil,
			Data: 0,
		},
	}
}

// 获取最后的节点
func (l *LinkList) LastNode() *Node {
	if l.Header == nil {
		return nil
	}
	var flag = l.Header
	var next = l.Header.Next
	for next != nil {
		flag = next
		next = next.Next
	}
	return flag
}

// 获取最后的节点
func TestLastNode(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})

	fmt.Println(l.LastNode())
}

// 获取链表长度
func (l *LinkList) Len() int {
	var next *Node
	next = l.Header.Next
	var count int
	for next != nil {
		count++
		next = next.Next
	}
	return count
}

// 获取链表长度
func TestLinkListLen(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})

	fmt.Print(l.Len())
}

func (l *LinkList) Insert(node *Node) {
	lastNode := l.LastNode()
	lastNode.Next = node
}

func (l *LinkList) Print() {
	var rs = make([]int, 0, 10)

	if l.Header == nil {
		fmt.Println(rs)
	}
	var next = l.Header.Next
	for next != nil {
		rs = append(rs, next.Data)
		next = next.Next
	}
	fmt.Println(rs)
}

func TestPrint(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})

	l.Print()
}

func TestReverseLinkList(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})
	l.Insert(&Node{nil, 8})
	l.Insert(&Node{nil, 9})
	l.Insert(&Node{nil, 10})

	var flag, next *Node
	next = l.Header.Next
	var tmp *Node
	for next!=nil {
		tmp = next.Next
		next.Next = flag

		flag=next
		next = tmp
	}
	l.Header.Next = flag
	l.Print()
}
