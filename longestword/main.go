package main

import (
	"fmt"
)

// Checks if the given word is a subsequence of the given base string.
func isSubsequence(base, word string) bool {
	var cursor int
	for _, c := range word {
		var matched bool
		for i := cursor; i < len(base); i++ {
			if byte(c) == base[i] {
				cursor = i + 1
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

// Given a base string, find the longest word in the input list of words which is
// a subsequence of the given base string.
//
// Runtime complexity is O(N * W * L), where
//   N = length if input base string
//   W = number of words in the dictionary
//   L = number of chars in the longest word in the dictionary
func findLongestSubsequence(base string, words []string) string {
	var best string
	for _, word := range words {
		if !isSubsequence(base, word) {
			continue
		}

		if len(best) < len(word) {
			best = word
		}
	}

	return best
}

// Finds the longest subsequence of an input string from a given list of words
// by simultaneously matching to all words.
//
// Runtime complexity is O(N * W * L), where
//   N = length if input base string
//   W = number of words in the dictionary
//   L = number of chars in the longest word in the dictionary
func findLongestSubsequenceV2(base string, words []string) string {
	wordsSubsequenceCounter := make([]int, len(words))
	matches := make([]bool, len(words))
	for _, c := range base {
		for i, word := range words {
			counter := wordsSubsequenceCounter[i]

			// If matching is already done, skip!
			if counter >= len(word) {
				continue
			}

			if byte(c) == word[counter] {
				wordsSubsequenceCounter[i] = counter + 1
			}
			if counter >= len(word)-1 {
				matches[i] = true
			}
		}
	}

	var longestWord string
	for i, word := range words {
		if !matches[i] {
			continue
		}

		if len(longestWord) < len(word) {
			longestWord = word
		}
	}

	return longestWord
}

type tuple struct {
	word    string
	matches int
}

func (t *tuple) done() bool {
	return t.matches == len(t.word)-1
}

func findLongestSubsequenceV3(base string, words []string) string {
	buckets := make(map[rune][]tuple)

	// TODO: Implement me!

	return ""
}

func main() {
	words := []string{"able", "ale", "apple", "bale", "kangaroo"}
	input := "abppplee"

	fmt.Println(findLongestSubsequence(input, words))
}
