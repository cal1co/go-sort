package sort

import "math/rand"

// Shuffle is a linear time shuffle operation that reliably reorganizes an integer array
//
// Time Complexity: O(N)
//
// Space: O(1)
func Shuffle(arr []int) {
	for i := 0; i < len(arr); i++ {
		idx := rand.Intn(len(arr))
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
