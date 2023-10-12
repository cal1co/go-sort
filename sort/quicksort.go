package sort

// Quicksort is a versatile sorting algorithm that uses a pivot element to partition
// a list into two sub-arrays, and sorting them recursively.
//
// Time Complexity: O(N log N) on average, O(N^2) worst-case
//
// Space: O(log N)
//
// Note: Quicksort's performance relies on pivot selection, and it is subject to
// variations in performance based on pivot choices.
func QuickSort[T Number](arr []T) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

func partition[T Number](arr []T) int {
	pivot := arr[len(arr)-1]
	i := -1

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}
