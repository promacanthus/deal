package bitmap

type BitMap struct {
	bytes []byte
	nbits int
}

func NewBitMap(bits int) *BitMap {
	return &BitMap{
		bytes: make([]byte, bits/8+1),
		nbits: bits,
	}
}

func (b *BitMap) set(k int) {
	if k > b.nbits {
		return
	}
	byteIndex := k / 8
	bitIndex := k % 8
	b.bytes[byteIndex] |= 1 << bitIndex
}

func (b *BitMap) get(k int) bool {
	if k > b.nbits {
		return false
	}
	byteIndex := k / 8
	bitIndex := k % 8
	return (b.bytes[byteIndex] & (1 << bitIndex)) != 0
}
