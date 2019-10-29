package main

import (
	"bytes"
	"fmt"
)

type node struct {
	value int
	next *node
}

type stack struct {
	top *node
}

func (s *stack) pop() *node {
	res := s.top
	if res != nil {
		s.top = res.next
	}
	return res
}

func (s *stack) push(n *node) {
	n.next = s.top
	s.top = n
}

func (s stack) String() string {
	var buf bytes.Buffer
	for cur := s.top; cur != nil; {
		buf.WriteString(fmt.Sprintf(": %d", cur.value))
		cur = cur.next
	}
	return buf.String()
}

func hanoi(items int, source, buffer, target *stack) {
	if items == 1 {
		target.push(source.pop())
	} else {
		hanoi(items - 1, source, target, buffer)
		hanoi(1, source, buffer, target)
		hanoi(items - 1, buffer, source, target)
	}
}

func main() {
	var source, buffer, target stack
	items := 20
	for i := 0; i < items; i++ {
		source.push(&node{value: i})
	}
	
	hanoi(items, &source, &buffer, &target)

	fmt.Println("source:", source)
	fmt.Println("buffer:", buffer)
	fmt.Println("target:", target)
}
