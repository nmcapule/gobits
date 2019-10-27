package main

import "fmt"

func encoder(words, codeWords []string) []string {
	table := make(map[string]string)
	nextAvailable := 0

	for i, word := range words {
		_, ok := table[word]
		if !ok {
			table[word] = codeWords[nextAvailable]
			nextAvailable += 1
		}

		words[i], _ = table[word]
	}

	return words
}

func main() {
	fmt.Println(encoder([]string{"a"}, []string{"1", "2", "3", "4"}))
	fmt.Println(encoder([]string{"a", "b"}, []string{"1", "2", "3", "4"}))
	fmt.Println(encoder([]string{"a", "b", "a"}, []string{"1", "2", "3", "4"}))
}
