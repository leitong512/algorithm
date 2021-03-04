package _7_linkedlist

import "testing"

var l *LinkedList
func init() {
	n5 := &ListNode{Value: 5}
	n4 := &ListNode{Value: 4, Next: n5}
	n3 := &ListNode{Value: 3, Next: n4}
	n2 := &ListNode{Value: 2, Next: n3}
	n1 := &ListNode{Value: 1, Next: n2}
	l = &LinkedList{Head: &ListNode{Next: n1}}   //head 哨兵
}
func TestLinkedList_Reverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestLinkedList_HasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.Head.Next.Next.Next.Next.Next.Next = l.Head.Next.Next
	t.Log(l.HasCycle())
}
