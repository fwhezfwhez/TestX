package test_practice

import "fmt"

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

// 获取无环链表长度
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

func (l *LinkList) Insert(node *Node) *LinkList {
	lastNode := l.LastNode()
	lastNode.Next = node
	return l
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
func (l *LinkList) HasRing() bool{
	var trace = make(map[*Node]struct{}, 0)
	var next = l.Header.Next
	for next != nil {
		trace[next] = struct{}{}

		next = next.Next

		if _, ok := trace[next]; ok {
			return true
		}
	}
	return false
}
