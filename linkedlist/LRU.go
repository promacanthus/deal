package linkedlist

type DLNode struct {
	key, value int
	prev, next *DLNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLNode
	head, tail *DLNode
}

func newDLNode(key, value int) *DLNode {
	return &DLNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLNode),
		head:     newDLNode(0, 0),
		tail:     newDLNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) addToHead(node *DLNode) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	// 从哈希表中判断，不存在返回-1
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	// 获取哈希表中存的节点
	node := this.cache[key]
	// 将该节点移动到链表头
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// 从哈希表中判断
	if _, ok := this.cache[key]; !ok {
		// 	如不存在，新建节点并添加到哈希表和链表头
		node := newDLNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		// 判断当前是否超过缓存容量
		if this.size > this.capacity {
			// 在链表和哈希表中删除最久未使用的
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		// 如果存在，直接更新节点中value的值
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}
