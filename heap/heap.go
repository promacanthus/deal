package heap

type Heap struct {
	items []int
	los   bool // true == large heap;false == small heap
}

func NewHeap(flag bool) *Heap {
	return &Heap{
		items: make([]int, 0),
		los:   flag,
	}
}

func InitHeap(values []int, flag bool) *Heap {
	res := NewHeap(flag)
	for i := 0; i < len(values); i++ {
		res.Push(values[i])
	}
	return res
}

func (h *Heap) len() int {
	return len(h.items)
}

func (h *Heap) less(i, j int) bool {
	if h.los {
		return h.items[i] > h.items[j]
	}
	return h.items[i] < h.items[j]
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) Push(v int) {
	h.items = append(h.items, v)
	up(h, h.len()-1)
}

func up(h *Heap, i int) {
	for {
		p := (i - 1) / 2 // parent
		if p == i || !h.less(i, p) {
			break
		}
		h.swap(p, i)
		i = p
	}
}

func (h *Heap) Pop() int {
	n := h.len() - 1
	h.swap(0, n)
	down(h, 0, n)
	res := h.items[n]
	h.items = h.items[:n]
	return res
}

func down(h *Heap, root int, n int) {
	i := root
	for {
		l := 2*i + 1 // left
		if l >= n || l < 0 {
			break
		}
		max := l
		r := l + 1 // right
		if r < n && h.less(r, l) {
			max = r
		}
		if !h.less(max, i) {
			break
		}
		h.swap(i, max)
		i = max
	}
}

func (h Heap) Sort() []int {
	n := h.len()
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, h.Pop())
	}
	return res
}
