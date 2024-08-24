package dsa

import (
	"fmt"
	"log"
)

type Element struct {
	value int
	next  *Element
}

type LL struct {
	Head *Element
}

func (l *LL) appendLL(value int) {
	newElemnt := Element{value: value}
	if l.Head == nil {
		l.Head = &newElemnt
		return
	}

	currentElement := l.Head
	for currentElement.next != nil {
		currentElement = currentElement.next
	}
	currentElement.next = &newElemnt
}

func (l *LL) PrintElements() {
	currentElement := l.Head
	for currentElement != nil {
		fmt.Print(currentElement.value, "->")
		currentElement = currentElement.next
	}
	fmt.Print("NIL")
	fmt.Println()
}

func (l *LL) updateElement(value int, position int) {
	var updatedElement = Element{value: value}
	if l.Head == nil {
		l.Head = &updatedElement
		return
	}
	counter := 1
	currElement := l.Head
	for counter != position-1 {
		currElement = currElement.next
		counter++
	}
	temp := currElement.next.next
	currElement.next = &updatedElement
	updatedElement.next = temp
}

func (l *LL) deleteElement(position int) {
	if l.Head == nil {
		log.Println("Empty LL")
		return
	}
	counter := 1
	currElement := l.Head
	for counter != position-1 {
		currElement = currElement.next
		counter++
	}
	currElement.next = currElement.next.next
}
