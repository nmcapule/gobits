package main

import (
	"bytes"
	"fmt"
)

func stringSplosion(input string) string {
	var buffer bytes.Buffer

	for cursor := 1; cursor <= len(input); cursor++ {
		buffer.WriteString(input[:cursor])
	}

	return buffer.String()
}

func main() {
	fmt.Println(stringSplosion("Code"))
	fmt.Println(stringSplosion("abc"))
	fmt.Println(stringSplosion("ab"))
}
