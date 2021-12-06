package search

func LinearSearch(arr []int, want int) (int, bool) {
	for k, v := range arr {
		if want == v {
			return k, true
		}
	}
	return -1, false
}
