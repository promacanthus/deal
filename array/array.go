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
func (a *Array) Insert(index uint, v int) error {
	if a.length == uint(cap(a.data)) {
		a.expand()
		return errors.New("fully arrays")
	}

	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	// 移动数组，空出位置
	// 方法一：循环
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	// 方法二：copy
	copy(a.data[index:], a.data[index+1:])
	a.data = a.data[:a.length-1]

	// 将待插入位置的数据放到数组末尾，空出位置
	// 保证数组有序，则不能用这个方法
	// a.data[a.length] = a.data[index]

	a.data[index] = v
	a.length++
	return nil
}

// 插入在数组末尾
func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.length, v)
}

//
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return -1, errors.New("out of index range")
	}

	v := a.data[index]
	// 移动数组，填上空出位置
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	// 将数组末尾的数据放到待插入位置
	// 保证数组有序，则不能用这个方法
	// a.data[index] = a.data[a.length-1]
	a.length--
	return v, nil
}

func (a *Array) expand() {
	b := make([]int, a.length*2)
	for i := uint(0); i < a.length; i++ {
		b[i] = a.data[i]
	}
	a.data = b
}

// 打印数组
func (a *Array) String() {
	var format string
	for i := uint(0); i < a.length; i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
