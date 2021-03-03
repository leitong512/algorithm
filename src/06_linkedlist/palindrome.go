package _6_linkedlist

/* 思路：快慢指针，利用栈将慢指针里面的数据放到栈中，然后在一个一个的拿出来比较
	这里有length

	时间复杂度：O(n)
	空间复杂度：O(n)
 */
func IsPalindrome1(l *LinkedList) bool {
	if l.length == 0 {
		return false
	}
	if l.length == 1 {
		return true
	}
	length := l.length
	cur := l.head
	s := make([]rune, 0)
	for i := uint(1);i <= length;i++ {
		cur = cur.next
		if i== (length/2+1) && length % 2 != 0 {   //是奇数个，跳过中间结点
			continue
		}
		if i <= length/2 {  //前部分
			s = append(s, cur.GetValue().(rune))
		} else {
			if (s[length-i]) != cur.GetValue().(rune){
				return false
			}
		}
	}
	return true
}
/*
思路2
找到链表中间结点，将前半部分转置，在从中间向左右遍历对比
 */
func IsPalindrome2(l *LinkedList) bool {
	length := l.length
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	//todo  翻转链表
	//cur := l.head.next  //第一个结点
	//next := l.head.next.next
	for i := uint(0); i <= length/2; i++ {  //遍历到上半部分

	}
	return true
}
