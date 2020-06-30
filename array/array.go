package array

import (
	"errors"
	"fmt"
)

// 数组中数据为int类型

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

// 获取数组长度
func (a *Array) Len() uint {
	return a.length
}

// 判断数组是否越界
func (a *Array) isIndexOutOfRange(index uint) bool {
	return index > uint(cap(a.data))
}

// 通过索引查找数组
func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}

	return a.data[index], nil
}

//  插入数据到索引index
func (a Array) Insert(index uint, v int) error {
	if a.length == uint(cap(a.data)) {
		return errors.New("fully array")
	}

	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// 插入在数组末尾
func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

//
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}

	v := a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

// 打印数组
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.length; i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
