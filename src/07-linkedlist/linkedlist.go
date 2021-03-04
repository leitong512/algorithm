package _7_linkedlist

import "fmt"

//结点
type ListNode struct {
	Next *ListNode
	Value interface{}
}

//单链表
type LinkedList struct {
	Head *ListNode
}

//打印链表
func (this *LinkedList)Print() {
	cur := this.Head.Next
	format := ""
	for cur !=nil {
		format += fmt.Sprintf("%+v", cur.Value)
		cur = cur.Next
		if cur !=nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
单链表反转
时间复杂度：O(n)
 */
func (this *LinkedList)Reverse() {
	if this.Head == nil || this.Head.Next == nil || this.Head.Next.Next == nil {
		return
	}
	cur := this.Head.Next
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	this.Head.Next = pre
}

/*
判断链表是否有环
 */
func (this *LinkedList)HasCycle() bool {
	if this.Head == nil {
		return false
	}
	slow := this.Head
	fast := this.Head
	for fast != nil && fast.Next !=nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
两个单链表合并
 */
