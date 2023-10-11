package main

import (
	"fmt"

	"github.com/cal1co/go-sort/sort"
)

func main() {
	arr := []int{1, 999, 222, 342, 14, 2}
	fmt.Println("BUBBLE")
	sort.BubbleSort(arr)
	fmt.Println(arr)

	shuffle(arr)

	fmt.Println("MERGE")
	sort.MergeSort(arr)
	fmt.Println(arr)

	shuffle(arr)

	fmt.Println("QUICK")
	sort.QuickSort(arr)
	fmt.Println(arr)
}

func shuffle(arr []int) {
	fmt.Printf("SHUFFLE\n")
	sort.Shuffle(arr)
	fmt.Println(arr)
}
