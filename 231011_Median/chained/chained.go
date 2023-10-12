package chained

import (
	"fmt"
	"time"
)

type element struct {
	val  int
	head *element //erstes Element der Liste! head.val enthält den Count der Liste!
	prev *element
	next *element
}

func Run(values []int) {
	fmt.Println("chained")
	start := time.Now()
	medians := ContinuousMedian(values)
	end := time.Now()
	if len(medians) < 20 {
		fmt.Println(" => medians: ", medians)
	}
	fmt.Println(" => needed:", end.Sub(start))
}

func initList() *element {
	//Erzeuge leere Liste
	list := &element{val: 0, next: nil, prev: nil}
	list.head = list

	return list
}

func printList(list *element, head string) {
	fmt.Print(head)

	//erstes Element suchen, da ich nicht weiß wo ich in der Liste stehe!
	for {
		if list.prev == nil {
			break
		}
		list = list.prev
	}

	for {
		//wenn es ein nächstes Element gibt
		if list.next == nil {
			break
		}
		//dann gehe zum nächsten Element und gib den Wert aus
		list = list.next
		fmt.Print(" -> ")
		fmt.Print(list.val)
	}
	fmt.Println()
}

func appendSortlist(list *element, val int) *element {
	if val >= list.val {
		//Wenn der Wert größer ist als der Wert des aktuellen Elements
		//...Dann prüfe das nächste Element
		for {
			if list.next == nil || val <= list.next.val {
				break
			}
			list = list.next
		}
	} else {
		//Wenn der Wert kleiner ist als der Wert des aktuellen Elements
		//Dann prüfe das vorherige element
		for {
			if list.prev == nil || val >= list.val {
				break
			}
			list = list.prev
		}
	}

	// Das Element soll hinter das aktuelle Element eingefügt werden
	newElement := &element{val: val, head: list.head, prev: list, next: list.next}
	//Neues Element eingefügt, Counter muss erhöht werden!
	list.head.val++
	// Der Nchfolger des aktuellen muss dann auf das neue als Vorgänger haben
	if list.next != nil {
		list.next.prev = newElement
	}
	// Der Nachfolger des aktuellen Elements ist das neue Element
	list.next = newElement

	return newElement
}

func ContinuousMedian(values []int) []float64 {
	// Partial das Array aufspalten um den Kontinuierlichen Median zu berechnen
	partitions := initList()
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	for i := 0; i < len(values); i++ {
		partitions = appendSortlist(partitions, values[i])
		medians = append(medians, Median(partitions))
	}

	return medians
}

func Median(list *element) float64 {
	list = list.head  //erstes Element der Liste!
	count := list.val //Anzahl der Elemente in der Liste

	//in die Mitte der Liste gehen
	for i := 0; i <= count/2; i++ {
		list = list.next
	}

	// Grade
	if count%2 == 0 { //Head.val enthält den Count der Liste!
		// Mittelwert der beiden mittleren Werte
		return float64(list.val+list.prev.val) / 2
	}

	// Ungrade
	return float64(list.val)
}
