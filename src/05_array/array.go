package _5_array

import (
	"errors"
	"fmt"
)

/*
	1) 数组的插入、删除、按照下标随机访问操作
	2）数组中的数据是 int 类型的
 */

type Array struct {
	data []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data :make([]int,capacity),
		length: 0,
	}
}

func (a *Array)Len() uint{
	return a.length
}

//判定数组是否越界
func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0, n-1]
func (a *Array)Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index],nil
}

//插入数值到索引index上
func (a *Array)Insert(index uint, v int) error {
	if a.length == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := a.length; i > index;i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

func (a *Array)InsertToTail(v int) error {
	return a.Insert(a.Len(),v)
}

//删除索引index上的值

func (a *Array)Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len() -1 ;i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

//打印数列
func (a *Array)Print() {
	var format string
	for i := uint(0);i < a.Len();i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
