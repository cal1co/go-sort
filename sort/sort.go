package sort

import "math/rand"

func Shuffle(arr []int) {
	for i := 0; i < len(arr); i++ {
		idx := rand.Intn(len(arr))
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
