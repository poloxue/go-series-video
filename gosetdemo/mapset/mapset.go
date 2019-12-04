package mapset

// 这里面 key 用 interface 的好处，集合中的元素可以是任何类型
type MapSet map[interface{}]struct{}

// 数据结构经常会遇到一些方法，比如，增删查，
// 除了一些数据结构相关的操作，还有一些集合特有的方法，比如交并差

func NewMapSet(elements ...interface{}) *MapSet {
	set := make(MapSet)
	for _, e := range elements {
		set.Add(e)
	}

	return &set
}

func (set *MapSet) Add(e interface{}) bool {
	if _, found := (*set)[e]; found {
		return false
	}

	(*set)[e] = struct{}{}
	return true
}

func (set *MapSet) Remove(e interface{}) {
	delete(*set, e)
}

func (set *MapSet) Contains(e interface{}) bool {
	_, found := (*set)[e]
	return found
}

func (set *MapSet) Cardinatity() int {
	return len(*set)
}

func (set *MapSet) Intersect(other *MapSet) *MapSet {
	newSet := NewMapSet()

	// 为了高效的进行交集操作，先判断执行集合运算的两个集合的长度
	if set.Cardinatity() < other.Cardinatity() {
		for _, e := range *set {
			if _, found := (*other)[e]; found {
				newSet.Add(e)
			}
		}
	} else {
		for _, e := range *other {
			if _, found := (*set)[e]; found {
				newSet.Add(e)
			}
		}
	}

	return newSet
}

func (set *MapSet) Union(other *MapSet) *MapSet {
	unionSet := new(MapSet)

	for _, e := range *set {
		unionSet.Add(e)
	}

	for _, e := range *other {
		unionSet.Add(e)
	}

	return unionSet
}

func (set *MapSet) Difference(other *MapSet) *MapSet {
	// set - other
	differenceSet := NewMapSet()
	for _, e := range *set {
		if _, found := (*other)[e]; !found {
			differenceSet.Add(e)
		}
	}

	return differenceSet
}
