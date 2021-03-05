package _8_stack

import "fmt"

type ListNode struct {
	Value interface{}
	Next *ListNode
}
type StackLinkedList struct {
	Head *ListNode
}

func NewLinkedListStack() *StackLinkedList {
	return &StackLinkedList{
		Head: &ListNode{},
	}
}

func (this *StackLinkedList)IsEmpty() bool {
	return this.Head.Next == nil
}
func (this *StackLinkedList)Push(v interface{}) {
	newNode := &ListNode{
		Value: v,
	}
	newNode.Next = this.Head.Next
	this.Head.Next = newNode
}
func (this *StackLinkedList)Pop() interface{} {
	if this.Head.Next == nil {
		return nil
	}
	v := this.Head.Next.Value
	this.Head.Next = this.Head.Next.Next
	return v
}
func (this *StackLinkedList)Top() interface{} {
	if this.Head.Next == nil {
		return nil
	}
	return this.Head.Next.Value
}
func (this *StackLinkedList)Flush() {
	this.Head.Next = nil
}
func (this *StackLinkedList)Print() {
	if this.Head.Next == nil {
		fmt.Println("empty stack")
		return
	}
	cur := this.Head.Next
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next
	}
}