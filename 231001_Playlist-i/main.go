package main

// 01.10.2023 - https://www.youtube.com/watch?v=kPR3IWC0Wvc&t=4s

//Aufgabenstellung:
// Entfene das i-te Element von hinten aus einer Liste
//
//Gegeben sind:
// Eine einfach verkettete Liste und eine Ganzzahl i.
//
//Beispiel:
// Liste: 1 -> 2 -> 3 -> 4 -> 5
// i: 2
// Ergebnis: 1 -> 2 -> 3 -> 5
//
//Hinweis:
// Die Lösung muss Single-Pass sein und mit einem konstanten Speiocherbedarf O(1) auskommen.
//
//Lösung:
// Zwei Zeiger, einer auf der 0te Element (Vorgänger) und einer auf das erste Element (letztes)
// Ein Zeiger wird um i Elemente nach vorne geschoben.
// Beide Zeiger werden gleichzeitig weiter geschoben, bis der erste Zeiger am Ende der Liste angekommen ist.
// Dann zeigt der andere Zeiger auf das vorherigen Element von dem entfernt werden soll.
// Das Element wird entfernt, indem der Zeiger des vorherigen Elements auf das Element nach dem zu entfernenden Element gesetzt wird!

import (
	"fmt"
)

// Struct mit einem Zeiger und einem Wert als verkettete Liste
type element struct {
	next *element
	val  int
}

func main() {
	playlist := initList(5)
	printList(playlist, "initial list:")

	removeElement(playlist, 2)
	printList(playlist, "new list:")

	fmt.Println("done!")
}

func initList(size int) *element {
	firstElement := &element{val: 0, next: nil}
	lastElement := firstElement

	for i := 1; i <= size; i++ {
		newElement := &element{val: i, next: nil}
		lastElement.next = newElement
		lastElement = newElement
	}

	return firstElement
}

func printList(list *element, head string) {
	fmt.Println(head)

	for {
		//0tes soll nicht mit ausgegeben werden... dowhile
		list = list.next
		if list == nil {
			break
		}

		fmt.Print(list.val)
		if list.next != nil {
			fmt.Print(" -> ")
		}
	}

	fmt.Println("")
}

func removeElement(list *element, i int) *element {
	// Zwei Zeiger, einer auf der 0te Element (Vorgänger) und einer auf das erste Element (letztes)
	previousElement := list
	lastElement := list.next

	// Ein Zeiger wird um i Elemente nach vorne geschoben.
	for j := 1; j < i; j++ {
		lastElement = lastElement.next
	}

	// Beide Zeiger werden gleichzeitig weiter geschoben, bis der erste Zeiger am Ende der Liste angekommen ist.
	for j := 0; lastElement.next != nil; j++ {
		previousElement = previousElement.next
		lastElement = lastElement.next
	}
	// Dann zeigt der andere Zeiger auf das vorherigen Element von dem entfernt werden soll.

	// Das Element wird entfernt, indem der Zeiger des vorherigen Elements auf das Element nach dem zu entfernenden Element gesetzt wird!#
	previousElement.next = previousElement.next.next

	return nil
}
