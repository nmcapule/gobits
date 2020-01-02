package main

import "fmt"

// Runtime complexity is O(NlogN)
// Space complexity is O(NlogN)
//
// Might be better with heap sort?
func uniqueSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var left, right []int
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == pivot {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	sorted := uniqueSort(left)
	sorted = append(sorted, pivot)
	sorted = append(sorted, uniqueSort(right)...)

	return sorted
}

func herpderp(arr []int, i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if arr[parent] < arr[i] {
		tmp := arr[parent]
		arr[parent] = arr[i]
		arr[i] = tmp
		herpderp(arr, parent)
	}
}

func heapify(arr []int) {
	for i := 1; i < len(arr); i++ {
		herpderp(arr, i)
	}
}

func heapSort(arr []int) {
	for limit := len(arr); limit > 0; limit-- {
		heapify(arr[:limit])

		tmp := arr[0]
		arr[0] = arr[limit-1]
		arr[limit-1] = tmp
	}
}

func main() {
	fmt.Println(uniqueSort([]int{}))
	fmt.Println(uniqueSort([]int{1}))
	fmt.Println(uniqueSort([]int{1, 1}))
	fmt.Println(uniqueSort([]int{1, 2, 1}))
	fmt.Println(uniqueSort([]int{1, 1, 2, 2, 4, 3, 2}))
}
