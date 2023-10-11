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
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	middle := len(arr) / 2
	left := make([]int, middle)
	right := make([]int, len(arr)-middle)
	copy(left, arr[:middle])
	copy(right, arr[middle:])

	MergeSort(left)
	MergeSort(right)

	merge(arr, left, right)
}

func merge(arr, left, right []int) {
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
