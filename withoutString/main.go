package main

import (
	"bytes"
	"fmt"
)

func withoutString(input string, exclude string) string {
	var buffer bytes.Buffer

	var i int
	for i < len(input) {
		if i+len(exclude) > len(input) {
			buffer.WriteString(input[i:])
			break
		}
		if input[i:i+len(exclude)] == exclude {
			i += len(exclude)
			continue
		}

		buffer.WriteByte(input[i])
		i += 1
	}

	return buffer.String()
}

func main() {
	fmt.Println(withoutString("Hello there", "llo"))
	fmt.Println(withoutString("Hello there", "e"))
	fmt.Println(withoutString("Hello there", "x"))
}
