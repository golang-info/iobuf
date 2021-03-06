// Copyright 2020 (C) Zhou Peng <p@iobuf.io>

package quicksort

import (
	"fmt"
)

func ExampleInts() {
	ints := []int{2, 1, 0, -1, -2}
	Ints(ints)
	fmt.Println(ints)

	// Output:
	// [-2 -1 0 1 2]
}

func ExampleFloat64s() {
	flts := []float64{2.2, 1.1, 0.0, -1.1, -2.2}
	Float64s(flts)
	fmt.Println(flts)

	// Output:
	// [-2.2 -1.1 0 1.1 2.2]
}

func ExampleStrings() {
	strs := []string{"the", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}
	Strings(strs)
	fmt.Println(strs)

	// Output:
	// [brown dog fox jumped lazy over quick the the]
}

type item struct {
	i int
	s string
}

type itemSlice []item

func (p itemSlice) Len() int { return len(p) }
func (p itemSlice) Less(i, j int) bool {
	if p[i].i < p[j].i {
		return true
	}
	if p[i].i == p[j].i {
		return p[i].s < p[j].s
	}
	return false
}
func (p itemSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func ExampleSort() {
	items := []item{
		item{2, "2"},
		item{1, "1"},
		item{0, "0"},
		item{-1, "-1"},
		item{-2, "-2"},
	}
	Sort(itemSlice(items))
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}

	// Output:
	// {-2 -2}
	// {-1 -1}
	// {0 0}
	// {1 1}
	// {2 2}
}
