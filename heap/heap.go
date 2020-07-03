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
	h.items[1], h.items[h.count] = h.items[h.count], h.items[1]
	h.count--
	res := heapify(h, h.count)
	return res, true
}

func heapify(h *Heap, count uint) int {
	for i := uint(1); i < count/2; {
		maxIndex := i
		if h.items[i] < h.items[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && h.items[maxIndex] < h.items[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		h.items[i], h.items[maxIndex] = h.items[maxIndex], h.items[i]
	}
	res := h.items[count]
	return res
}

func (h *Heap) Sort() {

	for i := uint(1); i <= h.count; i++ {
		h.Delete()
	}
}
