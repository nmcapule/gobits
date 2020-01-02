// Good morning! Here's your coding interview problem for today.
// This problem was asked by Google.
//
// An XOR linked list is a more memory efficient doubly linked list. 
// Instead of each node holding next and prev fields, it holds a field
// named both, which is an XOR of the next node and the previous node.
// Implement an XOR linked list; it has an add(element) which adds the
// element to the end, and a get(index) which returns the node at index.
//
// If using a language that has no pointers (such as Python), you can
// assume you have access to get_pointer and dereference_pointer functions
// that converts between nodes and memory addresses.
package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

type XorNode struct {
	Value int
	Both uintptr
}

type XorLinkedList struct {
	Head *XorNode
	Tail *XorNode
}

func NewXorLinkedList() *XorLinkedList {
	return &XorLinkedList{}
}

func (l *XorLinkedList) Add(value int) {
	node := &XorNode{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	
	l.Tail.Both ^= uintptr(unsafe.Pointer(node))
	node.Both ^= uintptr(unsafe.Pointer(l.Tail))
	l.Tail = node
}

func (l *XorLinkedList) Get(index int) *XorNode {
	var prev uintptr
	cursor := uintptr(unsafe.Pointer(l.Head))
	for i := 0; i < index; i++ {
		if prev == uintptr(unsafe.Pointer(l.Tail)) {
			return nil
		}
		
		node := (*XorNode)(unsafe.Pointer(cursor))
	
		tmp := cursor
		cursor = prev ^ node.Both
		prev = tmp
	}
	
	return (*XorNode)(unsafe.Pointer(cursor))
}

func (l *XorLinkedList) String() string {
	var buf bytes.Buffer
	var prev uintptr
	cursor := uintptr(unsafe.Pointer(l.Head))
	for {
		node := (*XorNode)(unsafe.Pointer(cursor))
		buf.WriteString(fmt.Sprintf(" > %d", node.Value))
	
		tmp := cursor
		cursor = prev ^ node.Both
		prev = tmp
		
		if node == l.Tail {
			break
		}
	}
	
	return buf.String()
}

func main() {
	list := NewXorLinkedList()
	list.Add(3)
	list.Add(2)
	list.Add(1)

	fmt.Println(list)
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
	fmt.Println(list.Get(3))
}
