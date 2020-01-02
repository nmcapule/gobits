package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func sumNumbers(input string) int64 {
	var sum int64

	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		i, _ := strconv.ParseInt(match, 10, 64)
		sum += i
	}

	return sum
}

func main() {
	fmt.Println(sumNumbers("abc123xyz"))
	fmt.Println(sumNumbers("aa11b33"))
	fmt.Println(sumNumbers("7 11"))
}
