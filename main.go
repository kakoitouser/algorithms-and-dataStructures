package main

import (
	"fmt"

	"github.com/kakoitouser/algorithms-and-dataStructures/search"
	"github.com/kakoitouser/algorithms-and-dataStructures/sort"
)

func main() {
	arr := []int{9, 8, 6, 5, 1, 0, 78, -36}
	sort.RunAlgWithTimer(arr, sort.BubbleSort)
	sort.RunAlgWithTimer(arr, sort.BubbleSort2)
	sort.RunAlgWithTimer(arr, sort.SimpleSort)

	///
	want := 78
	if ind, ok := search.LinearSearch(arr, want); ok {
		fmt.Printf("the index of %d in array = %d\n", want, ind)
	}
}
