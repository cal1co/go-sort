package sort

// Mergesort is a divide-and-conquer sorting algorithm that efficiently sorts a list
// by recursively dividing it into smaller sub-arrays, sorting them, and then merging
// the sorted sub-arrays.
//
// Time Complexity: O(N log N)
//
// Space: O(N)
//
// Note: Mergesort guarantees a stable and efficient sorting approach, making it
// suitable for various scenarios, including large datasets.
func MergeSort[T Number](arr []T) {
	if len(arr) <= 1 {
		return
	}

	middle := len(arr) / 2
	left := make([]T, middle)
	right := make([]T, len(arr)-middle)
	copy(left, arr[:middle])
	copy(right, arr[middle:])

	MergeSort(left)
	MergeSort(right)

	merge(arr, left, right)
}

func merge[T Number](arr, left, right []T) {
	l, r := 0, 0
	for i := 0; i < len(arr); i++ {
		if l < len(left) && (r >= len(right) || left[l] <= right[r]) {
			arr[i] = left[l]
			l++
		} else {
			arr[i] = right[r]
			r++
		}
	}
}
