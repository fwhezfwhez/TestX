package test_practice

import (
	"fmt"
	"testing"
)


// 获取最后的节点
func TestLastNode(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})

	fmt.Println(l.LastNode())
}


// 获取链表长度
func TestLinkListLen(t *testing.T) {
	l := NewLinkList()
	l.Insert(&Node{nil, 5})
	l.Insert(&Node{nil, 6})
	l.Insert(&Node{nil, 7})

	fmt.Print(l.Len())
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
	for next != nil {
		tmp = next.Next
		next.Next = flag

		flag = next
		next = tmp
	}
	l.Header.Next = flag
	l.Print()
}
