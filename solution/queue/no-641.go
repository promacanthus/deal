package queue

type MyCircularDeque struct {
	values               []int
	head, tail, capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		values:   make([]int, k),
		head:     0,
		tail:     0,
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.head == 0 {
		if this.tail == this.capacity-1 {
			return false
		} else {
			this.values[this.capacity-1] = value
			this.head = this.capacity - 1
			return true
		}
	}

	if this.head == this.tail {
		return false
	} else {
		this.values[this.head-1] = value
		this.head--
		return true
	}
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if (this.tail+1)%this.capacity == this.head {
		return false
	}
	this.values[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == this.tail {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.head == this.tail {
		return false
	}
	this.tail = (this.tail - 1) % this.capacity
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.head == this.tail {
		return -1
	}
	return this.values[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.head == this.tail {
		return -1
	}
	return this.values[this.tail]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.tail+1)%this.capacity == this.head {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
