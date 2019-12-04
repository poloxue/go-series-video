package bitset

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	args := []int{1, 2, 3, 8, 100, 99}
	set := NewBitSet(args[:len(args)-1]...)
	for _, n := range args {
		if set.Contains(n) {
			fmt.Printf("%d 存在\n", n)
		} else {
			fmt.Printf("%d 不存在\n", n)
		}
	}

	set.Add(99)
	if set.Contains(99) {
		fmt.Println("99 存在")
	}

	set.Clear(100)
	if set.Contains(100) {
		fmt.Println("100 存在")
	} else {
		fmt.Println("100 不存在")
	}
}

func TestIntersect(t *testing.T) {
	set1 := NewBitSet(1, 2, 3)
	set2 := NewBitSet(1, 4, 3)

	set3 := set1.Intersect(set2)
	fmt.Println(set3.data)
}

func TestUnion(t *testing.T) {
	set := NewBitSet(1, 2, 10, 99)
	set.Visit(func(n int) bool {
		fmt.Println(n)
		return false
	})
}
