package bitset

import "math/bits"

// 状态标志位 元素个数固定 一个整数的就足以表示左右的状态
// 集合场景 元素个数不确定 一个整数不足以表示
type BitSet struct {
	data []uint64
	size int // 用于存放集合元素的个数
}

func NewBitSet(ns ...int) *BitSet {
	if len(ns) == 0 {
		return new(BitSet)
	}

	// 确定给 bitset 分配多大存放集合元素
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}

	if max < 0 {
		return new(BitSet)
	}

	s := &BitSet{
		data: make([]uint64, index(max)+1),
	}

	// 添加元素
	for _, n := range ns {
		if n >= 0 {
			s.data[index(n)] |= posVal(n)
			s.size++
		}
	}

	return s
}

func (set *BitSet) Contains(n int) bool {
	// 应该存在什么位置
	i := index(n)
	if i >= len(set.data) {
		return false
	}

	return set.data[i]&posVal(n) != 0
}

func (set *BitSet) Clear(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		return set
	}

	if set.data[i]&posVal(n) != 0 {
		set.data[i] &^= posVal(n)
		set.size--
	}

	return set
}

func (set *BitSet) Add(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		ndata := make([]uint64, i+1)
		copy(ndata, set.data)
		set.data = ndata
	}

	if set.data[i]&posVal(n) == 0 {
		set.data[i] |= posVal(n)
	}

	return set
}

func (set *BitSet) Size() int {
	return set.size
}

func (set *BitSet) Intersect(other *BitSet) *BitSet {
	minLen := min(len(set.data), len(other.data))

	intersectSet := &BitSet{
		data: make([]uint64, minLen),
	}

	for i := 0; i < minLen; i++ {
		intersectSet.data[i] = set.data[i] & other.data[i]
	}

	intersectSet.size = set.computeSize()

	return intersectSet
}

func (set *BitSet) Union(other *BitSet) *BitSet {
	var maxSet, minSet *BitSet
	if len(set.data) > len(other.data) {
		maxSet, minSet = set, other
	} else {
		maxSet, minSet = other, set
	}

	unionSet := &BitSet{
		data: make([]uint64, len(maxSet.data)),
	}

	minLen := len(minSet.data)
	copy(unionSet.data[minLen:], maxSet.data[minLen:])

	for i := 0; i < minLen; i++ {
		unionSet.data[i] = set.data[i] | other.data[i]
	}
	unionSet.size = unionSet.computeSize()
	return unionSet
}

func (set *BitSet) Difference(other *BitSet) *BitSet {
	setLen := len(set.data)
	otherLen := len(other.data)

	differenceSet := &BitSet{
		data: make([]uint64, setLen),
	}

	minLen := setLen
	if setLen > otherLen {
		copy(differenceSet.data[otherLen:], set.data[otherLen:])
		minLen = otherLen
	}

	for i := 0; i < minLen; i++ {
		differenceSet.data[i] = set.data[i] &^ other.data[i]
	}

	differenceSet.size = differenceSet.computeSize()

	return differenceSet
}

func (set *BitSet) Visit(do func(int) (skip bool)) (aborted bool) {
	d := set.data
	for i, len := 0, len(d); i < len; i++ {
		w := d[i]
		if w == 0 {
			continue
		}

		n := i << shift // 0 << 6，还是 0，1 << 6 就是 64，2 << 6 的就是 128
		for w != 0 {
			// 000.....000100 64~128 的话，表示 66，即 64 + 2，这个 2 可以由结尾 0 的个数确定
			// 那怎么获取结果 0 的个数呢？可以使用 bits.TrailingZeros64 函数
			b := bits.TrailingZeros64(w)
			if do(n + b) {
				return true
			}
			w &^= 1 << uint64(b) // 将已经检查的位清零
		}
	}
	return false
}

func (set *BitSet) computeSize() int {
	d := set.data
	n := 0
	for i, l := 0, len(d); i < l; i++ {
		if w := d[i]; w != 0 {
			n += bits.OnesCount64(w)
		}
	}

	return n
}
