package _7_linkedlist

import "testing"

var l *LinkedList
func init() {
	n6 := &ListNode{Value: 6}
	n5 := &ListNode{Value: 5,Next: n6}
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

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{Value: 9}
	n4 := &ListNode{Value: 7, Next: n5}
	n3 := &ListNode{Value: 5, Next: n4}
	n2 := &ListNode{Value: 3, Next: n3}
	n1 := &ListNode{Value: 1, Next: n2}
	l1 := &LinkedList{Head: &ListNode{Next: n1}}
	l1.Print()
	n10 := &ListNode{Value: 10}
	n9 := &ListNode{Value: 8, Next: n10}
	n8 := &ListNode{Value: 6, Next: n9}
	n7 := &ListNode{Value: 4, Next: n8}
	n6 := &ListNode{Value: 2, Next: n7}
	l2 := &LinkedList{Head: &ListNode{Next: n6}}
	l2.Print()
	MergeSortedList(l1, l2).Print()
}
func TestLinkedList_DeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(1)
	l.Print()
}
func TestLinkedList_FindMiddleNode(t *testing.T) {
	l.Print()
	t.Log(l.FindMiddleNode())
}
