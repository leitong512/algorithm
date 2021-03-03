package _6_linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", "上海自來水來自海上"}
	for _, str := range strs {
		l := NewLinkedList()
		for _, c:= range str {
			l.InsertToTail(c)
		}
		l.Print()
		t.Log(IsPalindrome1(l))
	}
}