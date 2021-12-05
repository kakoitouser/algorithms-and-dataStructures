package main

import (
	"fmt"

	"github.com/kakoitouser/algorithms-and-dataStructures/sort"
)

func main() {
	arr := []int{9, 8, 6, 5, 1, 0, 78, -36}
	fmt.Println(sort.BubbleSort(arr))
	fmt.Println(sort.BubbleSort([]int{}))
}
