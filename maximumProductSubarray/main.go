package main

import "fmt"

func imaxmin(xs ...int) (int, int) {
	a := xs[0]
	b := xs[0]
	for _, v := range xs {
		if v > a {
			a = v
		}
		if v < b {
			b = v
		}
	}
	return a, b
}

func maxProduct(nums []int) int {
	accmax := 0
	accmin := 0
	max := 0
	for i, v := range nums {
		if i == 0 {
			accmax = v
			accmin = v
			max = v
			continue
		}

		accmax, accmin = imaxmin(accmax*v, accmin*v, v)
		max, _ = imaxmin(max, accmax)
	}

	return max
}

func main() {
	fmt.Println(maxProduct([]int{-1, -2, -9, -6}))
}
