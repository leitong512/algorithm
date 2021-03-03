package _6_linkedlist

import "fmt"

/*
	单链表基本操作：查找，删除，插入，新建
 */

type ListNode struct {
	value interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	length uint
}

func NewListNode( v interface{}) *ListNode {
	return &ListNode{
		value: v,
		next: nil,
	}
}
func (l *ListNode)GetNext() *ListNode {
	return l.next
}
func (l *ListNode)GetValue() interface{} {
	return l.value
}
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: NewListNode(0),
		length: 0,
	}
}

//在某个结点后面插入
func (this *LinkedList)InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	newNode.next = p.next
	p.next = newNode
	this.length++
	return true
}
//在某个结点前面插入
func (this *LinkedList)InsertBefore(p *ListNode,  v interface{}) bool {
	if p == nil {
		return false
	}
	//找到 p 的前一个结点
	tmp := this.head
	for tmp.next != nil {
		if tmp.next == p {
			break
		}
		tmp = tmp.next
	}
	if tmp.next == nil { //没有找到
		return false
	}
	newNode := NewListNode(v)
	newNode.next = tmp.next
	tmp.next = newNode
	this.length++
	return true
}

//在链表头部插入结点
func (this *LinkedList)InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head,v)
}

//在链表尾部插入结点
func (this *LinkedList)InsertToTail(v interface{}) bool {
	tmp := this.head
	for tmp.next !=nil {
		tmp = tmp.next
	}
	return this.InsertAfter(tmp,v)
}
//通过索引查找结点  (index 从 0 开始)
func (this *LinkedList)FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}
//删除传入的结点
func (this *LinkedList)DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	del := this.head
	for cur != nil {
		if cur == p {
			break
		}
		cur = cur.next
		del = del.next
	}
	if cur == nil {   //没有找到
		return false
	}
	del.next = cur.next
	p = nil
	this.length--
	return true
}


func (this *LinkedList)Print() {
	tmp := this.head.next
	format := ""
	for tmp != nil {
		format += fmt.Sprintf("%+v", tmp.GetValue())
		tmp = tmp.next
		if tmp != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}