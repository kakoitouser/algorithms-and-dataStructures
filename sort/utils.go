package sort

import (
	"fmt"
	"time"
)

func calculateExecutionTime(arr []int, sortAlg func([]int) []int) ([]int, time.Duration) {
	start := time.Now()
	res := sortAlg(arr)
	end := time.Since(start)
	return res, end
}

func RunAlgWithTimer(arr []int, sortAlg func([]int) []int) {
	res, time := calculateExecutionTime(arr, sortAlg)
	fmt.Printf("sorted slice = %v, time=%v\n", res, time)
}
