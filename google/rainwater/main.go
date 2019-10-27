// https://techdevguide.withgoogle.com/paths/advanced/volume-of-water/#code-challenge
package main

import "fmt"

func rainwater(heightMap []int) int {
	leftMaxes := make([]int, len(heightMap))
	rightMaxes := make([]int, len(heightMap))

	max := 0
	for i, x := range heightMap {
		if x > max {
			max = x
		}
		leftMaxes[i] = max
	}

	max = 0
	for i := len(heightMap) - 1; i >= 0; i-- {
		x := heightMap[i]
		if x > max {
			max = x
		}
		rightMaxes[i] = max
	}

	volume := 0
	for i := 0; i < len(heightMap); i++ {
		min := leftMaxes[i]
		if rightMaxes[i] < min {
			min = rightMaxes[i]
		}

		volume += min - heightMap[i]
		heightMap[i] = min
	}

	return volume
}

func main() {
	fmt.Println(rainwater([]int{1, 3, 2, 4, 1, 3, 1, 4, 5, 2, 2, 1, 4, 2, 2}))
}
