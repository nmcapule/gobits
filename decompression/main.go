package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

type token struct {
	kind  string
	value string
}

func tokenize(input string) []token {
	var tokens []token

	var buffer bytes.Buffer
	var prevKind string

	for _, r := range input {
		var kind string
		if unicode.IsDigit(r) {
			kind = "NUMBER"
		} else if r == '[' {
			kind = "OPEN"
		} else if r == ']' {
			kind = "CLOSE"
		} else if unicode.IsLetter(r) {
			kind = "LETTER"
		} else {
			continue
		}

		if (kind != prevKind || kind == "OPEN" || kind == "CLOSE") && buffer.Len() > 0 {
			tokens = append(tokens, token{prevKind, buffer.String()})
			buffer.Reset()
		}

		buffer.WriteRune(r)
		prevKind = kind
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, token{prevKind, buffer.String()})
	}

	return tokens
}

func pushToStack(stack []token, tok token) []token {
	switch tok.kind {
	case "LETTER":
		if len(stack) > 0 && stack[len(stack)-1].kind == "LETTER" {
			tok.value = stack[len(stack)-1].value + tok.value
			stack = stack[:len(stack)-1]
			stack = pushToStack(stack, tok)
			break
		}
		fallthrough
	case "NUMBER":
		stack = append(stack, tok)
	case "CLOSE":
		multiplier, _ := strconv.Atoi(stack[len(stack)-2].value)
		letters := stack[len(stack)-1].value
		stack = stack[:len(stack)-2]

		var buffer bytes.Buffer
		for i := 0; i < multiplier; i++ {
			buffer.WriteString(letters)
		}

		stack = pushToStack(stack, token{"LETTER", buffer.String()})
	}
	return stack
}

func decompress(input string) []token {
	var stack []token
	tokens := tokenize(input)
	for _, tok := range tokens {
		stack = pushToStack(stack[:], tok)
	}
	return stack
}

func main() {
	fmt.Println(decompress("3[abc]4[ab]c"))
	fmt.Println(decompress("10[a]"))
	fmt.Println(decompress("2[3[a]b]"))
	fmt.Println(decompress("10[2[b1[xx]]]b"))
}
