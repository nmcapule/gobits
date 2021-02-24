package main

import (
	"bytes"
	"fmt"
)

func reverseWords(s string) string {
	// Let's do this the hard way.

	var buf bytes.Buffer
	marker := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if marker-i > 1 {
				if buf.Len() > 0 {
					buf.WriteString(" ")
				}
				buf.WriteString(s[i+1 : marker])
			}
			marker = i
		}
	}
	if marker > 0 {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(s[:marker])
	}

	return buf.String()
}

func main() {
	fmt.Println(reverseWords("a the sky is blue"))
}
