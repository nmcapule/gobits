// https://leetcode.com/problems/text-justification/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func justify(words []string, charsRequired, maxWidth int, left bool) string {
	spacesRequired := len(words) - 1
	spaces := maxWidth - (charsRequired - spacesRequired)

	var buf bytes.Buffer

	if spacesRequired == 0 || left {
		buf.WriteString(strings.Join(words, " "))
		for i := 0; i < maxWidth-charsRequired; i++ {
			buf.WriteRune(' ')
		}
		return buf.String()
	}

	spacesBetween := spaces / spacesRequired
	extras := spaces % spacesRequired

	for _, word := range words {
		buf.WriteString(word)

		budget := spaces
		if budget >= spacesBetween {
			budget = spacesBetween
		}
		if extras > 0 {
			budget += 1
			extras -= 1
		}
		spaces -= budget

		for i := 0; i < budget; i++ {
			buf.WriteRune(' ')
		}
	}

	return strings.TrimSpace(buf.String())
}

func fullJustify(words []string, maxWidth int) []string {
	var results []string
	var consumed, start int
	for cursor, word := range words {
		delta := len(word)
		if consumed > 0 {
			delta += 1
		}
		if consumed+delta <= maxWidth {
			consumed += delta
			continue
		}

		results = append(results, justify(words[start:cursor], consumed, maxWidth, false))

		consumed = len(word)
		start = cursor
	}
	return append(results, justify(words[start:], consumed, maxWidth, true))
}

func main() {
	words := []string{
		//"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		//"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do", "sd",
		"ask", "not", "what", "your", "country", "can",
		"do", "for", "you", "ask", "what", "you",
		"can", "do", "for", "your", "country",
	}

	lines := fullJustify(words, 16)
	for _, line := range lines {
		fmt.Println(line)
	}
}
