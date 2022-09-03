package main

import (
	"fmt"
	"sort"
)

func ExecSorting() {
	str := []string{"c", "a", "b"}
	sort.Strings(str)
	fmt.Println(str)

	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
