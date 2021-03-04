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
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.Head == nil || l1.Head.Next == nil {
		return l2
	}
	if l2 == nil || l2.Head == nil || l2.Head.Next == nil {
		return l1
	}
	l := &LinkedList{ Head: &ListNode{}}
	cur := l.Head
	cur1 := l1.Head.Next
	cur2 := l2.Head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.Value.(int) > cur2.Value.(int) {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}
	return l
}

/*
删除倒数第N个结点
 */

func (this *LinkedList)DeleteBottomN(n int) {
	if n <= 0 || this.Head == nil || this.Head.Next == nil {
		return
	}
	cur := this.Head
	for i := 0; i < n && cur != nil;i++ {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	del := this.Head
	for cur.Next != nil {
		cur = cur.Next
		del = del.Next
	}
	del.Next = del.Next.Next
}

/*
获取中间结点 (876)
 */
func (this *LinkedList)FindMiddleNode() *ListNode {
	if this.Head == nil || this.Head.Next == nil {
		return nil
	}
	fast := this.Head.Next
	slow := this.Head.Next
	for fast != nil && fast.Next != nil {
		fast =fast.Next.Next
		slow = slow.Next
	}
	return slow
}