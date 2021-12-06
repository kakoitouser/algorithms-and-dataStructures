package sort

// Сложность: O(n^2)
func SimpleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				// swap arr[i] and arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
