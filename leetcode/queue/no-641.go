package queue

type MyCircularDeque struct {
	// 空：head==tail
	// 满：(tail+1)%capacity==head
	values   []int
	capacity int

	// 指向第一个存放元素的位置
	// 插入时，索引减1，再赋值（注意取模）
	// 删除时，索引加1（注意取模）
	head int

	// 指向下一个存放元素的位置
	// 插入时，先赋值，再索引加1（注意取模）
	// 删除时，索引减1（注意取模）
	tail int

	// 循环队列
	// 加1的取模公式：（index + 1）% capacity
	// 减1的取模公式：（index - 1 + capacity）% capacity
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		values:   make([]int, k+1),
		head:     0,
		tail:     0,
		capacity: k + 1,
		// 	防止空和满状态判断的混淆，浪费1个存储单元
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + this.capacity) % this.capacity
	this.values[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.values[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + this.capacity) % this.capacity
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[(this.tail-1+this.capacity)%this.capacity]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.capacity == this.head
}
