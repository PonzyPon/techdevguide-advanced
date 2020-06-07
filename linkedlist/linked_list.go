package linkedlist

import (
	"fmt"
)

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type LinkedList struct {
	firstNode *Node
	length    int
}

type IndexOutOfBoundsError struct{}

func (ioobe IndexOutOfBoundsError) Error() string {
	return "IndexOutOfBoundsError"
}

func New() *LinkedList {
	l := LinkedList{
		length: 0,
	}
	return &l
}

func (l *LinkedList) first() (interface{}, error) {
	if l.length == 0 {
		return nil, IndexOutOfBoundsError{}
	}
	return l.first, nil
}

func (l *LinkedList) add(value interface{}) {
	if l.length == 0 {
		newNode := &Node{
			value: value,
		}
		l.firstNode = newNode
		l.length++
		return
	}
	lastNode := l.firstNode
	for i := 0; i < l.length-1; i++ {
		lastNode = lastNode.next
	}
	newNode := &Node{
		value: value,
		prev:  lastNode,
	}
	lastNode.next = newNode
	l.length++
}

func (l *LinkedList) toString() string {
	str := "["
	for node := l.firstNode; node != nil; node = node.next {
		str += fmt.Sprintf("%v, ", node.value)
	}
	str += "]"
	return str
}
