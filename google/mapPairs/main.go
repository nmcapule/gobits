package main

import "fmt"

func pairs(words []string) map[string]string {
	table := make(map[string]string)

	for _, word := range words {
		k := word[:1]
		v := word[len(word)-1:]
		table[k] = v
	}

	return table
}

func main() {
	fmt.Println(pairs([]string{"code", "bug"}))                  // → {"b": "g", "c": "e"}
	fmt.Println(pairs([]string{"man", "moon", "main"}))          // → {"m": "n"}
	fmt.Println(pairs([]string{"man", "moon", "good", "night"})) // → {"g": "d", "m": "n", "n": "t"}
}
