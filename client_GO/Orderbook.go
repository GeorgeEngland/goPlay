package main

import "strconv"

type Orderbook struct {
	bidList *List
	askList *List
}

type List struct {
	head *Element
	len  int
}

type Element struct {
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}

func (l *List) Init() *List {
	l.head.next = nil
	l.head.prev = nil
	l.len = 0
	return l
}

func (l *List) Len() int { return l.len }

func (l *List) Head() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insertAfter(e *Element, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) append(e *Element) *Element {
	e.prev = l.head
	e.next = nil
	e.list = l

	if l.Head() == nil {
		l.head = e
		e.prev = nil
		l.len++
		return e
	}
	currElem := l.Head()
	for currElem.next != nil {

		currElem = currElem.next
	}
	l.len++
	e.prev = currElem
	e.prev.next = e

	return e
}

func (l *List) String() string {
	if l.Head() == nil {
		return "Empty List"
	}
	return "list length: " + strconv.Itoa(l.Len())

}
