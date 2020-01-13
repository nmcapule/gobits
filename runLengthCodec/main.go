// Run-length encoding is a fast and simple method of encoding strings. The 
// basic idea is to represent repeated successive characters as a single count 
// and character. For example, the string "AAAABBBCCDAA" would be encoded as 
// "4A3B2C1D2A".
//
// Implement run-length encoding and decoding. You can assume the string to 
// be encoded have no digits and consists solely of alphabetic characters. You 
// can assume the string to be decoded is valid.
//
// We will be sending the solution tomorrow, along with tomorrow's question. 
// As always, feel free to shoot us an email if there's anything we can help with.
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func encode(s string) string {
	var buf bytes.Buffer
	var prev byte
	var count int

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] != prev {
			if count > 0 {
				buf.WriteString(fmt.Sprintf("%d%c", count, prev))
			}
			if i == len(s) {
				break
			}
			count = 1
		}
		if prev == s[i] {
			count += 1
		}
		prev = s[i]
	}

	return buf.String()
}

func decode(s string) string {
	var buf bytes.Buffer
	var count int

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			count = count*10 + int(s[i]-'0')
		} else {
			for j := 0; j < count; j++ {
				buf.WriteByte(s[i])
			}
			count = 0
		}
	}

	return buf.String()
}

func main() {
	fmt.Println(decode(encode("AAAAAAAAAAAAAAAaAABBBCCDAA")))
}
