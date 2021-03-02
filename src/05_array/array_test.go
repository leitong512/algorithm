package _5_array

import "testing"

func TestInsert(t *testing.T) {
	var capacity = 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity - 2;i++ {
		err := arr.Insert(uint(i),i+1)
		if err !=nil {
			t.Fatal(err)
		}
	}
	arr.Print()
	arr.Insert(uint(6),88)
	arr.Print()
	arr.InsertToTail(666)
	arr.Print()
}

func TestArray_Delete(t *testing.T) {
	var capacity = 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity ;i++ {
		err := arr.Insert(uint(i),i+1)
		if err !=nil {
			t.Fatal(err)
		}
	}
	arr.Print()
	for i := 9 ;i >=0;i -- {
		_, err := arr.Delete(uint(i))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestArray_Find(t *testing.T) {
	var capacity = 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity ;i++ {
		err := arr.Insert(uint(i),i+1)
		if err !=nil {
			t.Fatal(err)
		}
	}
	arr.Print()

	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}