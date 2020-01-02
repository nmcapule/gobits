// https://codingbat.com/prob/p183562
package main

import "fmt"

func makeBricks(small, big, goal int) bool {
	max := goal / 5
	if big < goal/5 {
		max = big
	}

	goal -= max * 5

	return goal <= small
}

func main() {
	fmt.Println(makeBricks(3, 1, 8))  // → true
	fmt.Println(makeBricks(3, 1, 9))  // → false
	fmt.Println(makeBricks(3, 2, 10)) // → true
}
