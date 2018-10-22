package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: fizzbuzz <number>")
		return
	}

	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid number: ", err)
		return
	}

	for i := 1; i <= max; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
