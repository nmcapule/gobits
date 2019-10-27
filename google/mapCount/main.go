package main

import "fmt"

func wordCount(words []string) map[string]int {
	table := make(map[string]int)

	for _, word := range words {
		v, _ := table[word]
		table[word] = v + 1
	}

	return table
}

func main() {
	fmt.Println(wordCount([]string{"a", "b", "a", "c", "b"}))
	fmt.Println(wordCount([]string{"c", "b", "a"}))
	fmt.Println(wordCount([]string{"c", "c", "c", "c"}))
}
