package main

import (
	"fmt"
	"strings"
)

// Node is a node from a single linked list
type Node struct {
	Data interface{}
	next *Node
}

// SingleList represents a single linked list
type SingleList struct {
	size int
	head *Node
}

// New returns an initialized SingleList.
func New() *SingleList {
	return &SingleList{}
}

// Size returns the number of nodes in the SingleList l
// The complexity is O(1).
func (l *SingleList) Size() int { return l.size }

// InsertFirst inserts a value v at the beginning of the SingleList l
func (l *SingleList) InsertFirst(v interface{}) {
	node := &Node{Data: v}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.size++
}

// RemoveFirst removes the first element of the SingleList l
func (l *SingleList) RemoveFirst() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}
	l.head = l.head.next
	l.size--
	return nil
}

// InsertLast inserts a value v at the end of the SingleList l
func (l *SingleList) InsertLast(v interface{}) {
	node := &Node{Data: v}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.size++
}

// RemoveLast removes the last element of the SingleList l
func (l *SingleList) RemoveLast() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}

	var prev *Node
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		l.head = nil
	}
	l.size--
	return nil
}

// Traverse traverses the SingleList l and prints all its values
func (l *SingleList) Traverse() string {
	if l.head == nil {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%v", l.head.Data))
	current := l.head.next
	for current != nil {
		sb.WriteString(fmt.Sprintf(", %v", current.Data))
		current = current.next
	}
	sb.WriteString("]")
	return sb.String()
}

// First returns the first element of the SingleList
func (l *SingleList) First() (interface{}, error) {
	if l.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return l.head.Data, nil
}
