package heap

type Heap struct {
	items []int
	count uint
}

func NewHeap() *Heap {
	items := make([]int, 1)
	items[0] = 0
	return &Heap{items: items, count: 0}
}

// max heap
func (h *Heap) Insert(v int) {
	h.items = append(h.items, v)
	h.count++
	update(h, v)
}

func update(heap *Heap, v int) {
	index := heap.count
	par := index / 2
	for v > heap.items[par] && par > 0 {
		heap.items[par], heap.items[index] = heap.items[index], heap.items[par]
		par, index = par/2, par
	}
}

func (h *Heap) Get() (int, bool) {
	if h.count == 0 {
		return -1, false
	}
	return h.items[1], true
}

func (h *Heap) Delete() (int, bool) {
	if h.count == 0 {
		return -1, false
	}
	res := heapify(h)
	return res, true
}

func heapify(h *Heap) int {

	h.items[1], h.items[h.count] = h.items[h.count], h.items[1]

	index := uint(1)
	leftChild := 2 * index
	rightChild := leftChild + 1

	for leftChild < h.count && rightChild < h.count {
		if h.items[index] < h.items[leftChild] {
			h.items[index], h.items[leftChild] = h.items[leftChild], h.items[index]
			index = leftChild
		} else if h.items[index] < h.items[rightChild] {
			h.items[index], h.items[rightChild] = h.items[rightChild], h.items[index]
			index = rightChild
		} else {
			break
		}
		leftChild = 2 * index
		rightChild = leftChild + 1
	}

	res := h.items[h.count]
	h.count--
	return res
}
