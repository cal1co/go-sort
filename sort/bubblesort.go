package sort

// BubbleSort repeatedly compares and swaps adjacent elements until the list is sorted.
//
// Time Complexity: O(N^2)
//
// Space: O(1)
//
// Note: While simple to understand, it's less efficient than alternatives like Quick Sort.
func BubbleSort[T Number](arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			a, b := arr[j], arr[j+1]
			if a > b {
				arr[j], arr[j+1] = b, a
			}
		}
	}
}
