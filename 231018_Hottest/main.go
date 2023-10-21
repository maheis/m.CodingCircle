package main

import (
	"container/list"
	"fmt"
	"time"
)

// 18.10.2023 - https://www.youtube.com/watch?v=0QPK9X1PtEE

// Aufgabenstellung:
// Gegeben sind eine Liste V und eine Ganzzahl k. Ermittle das laufende maximum für die Fenstergröße k.

// Beispiel:
// V = [ 27, 9, 17, 2, 12, 8 ]
// k = 3
// => [ 27, 17, 17, 12 ]

// Bedingungen:
// Die Laufzeit soll O(n) betragen, der Speicherbedarf soll O(k) betragen.
//
//Hinweis:
// Ein Algorithmus mit einer Laufzeit von O(n) hat eine lineare Laufzeit, die proportional zur Größe
// der Eingabe ist. Das bedeutet, dass die Laufzeit des Algorithmus linear mit der Größe der Eingabe wächst.
// Wenn die Größe der Eingabe um den Faktor k erhöht wird, sollte die Laufzeit des Algorithmus auch um den
// Faktor k erhöht werden.
// Der Speicherbedarf von O(k) bedeutet, dass der Speicherbedarf des Algorithmus proportional zur Größe
// des Eingabeparameters k ist. Mit anderen Worten, wenn die Größe des Eingabeparameters um den Faktor k
// erhöht wird, sollte der Speicherbedarf des Algorithmus auch um den Faktor k erhöht werden.
//
//Lösung:
// Es wird das Fenster betrachtet nicht die Werte. Dazu wird eine Que verwendet, die die Indizes der
// Elemente enthält. Die Que ist sortiert, so dass das erste Element immer das größte Element des aktuellen Fensters ist.
// Das nachfolgende Element enthält das größte Element für das Nachfolgende Fenster.
// Wenn das Fenster sich bewegt, wird das erste Element entfernt, wenn es aus dem Fenster läuft.
// Wenn ein neues Element hinzugefügt wird, wird das letzte Element entfernt, wenn es kleiner als das neue Element ist.
// Die Liste enthält maximal k Elemente, so dass der Speicherbedarf O(k) ist.
// Die Liste wird maximal n mal durchlaufen, so dass die Laufzeit O(n) ist.
//

func main() {
	// v := []int{27, 9, 17, 2, 12, 8}
	// 		   27, 9, 17
	// 			   9, 17, 2
	//				  17, 2, 12
	// 					  2, 12, 8
	// => [ 27 17 17 12 ]
	v := []int{27, 9, 17, 2, 12, 8, 9, 10, 25, 29, 3, 8, 9, 8, 10}
	// => [27 17 17 12 12 10 25 29 29 29 9 9 10]
	k := 3

	result := make([]int, 0, len(v)-k+1)
	var que list.List

	start := time.Now()
	//Itteriere durch die Liste...
	for i, n := range v {
		// Wenn der Zeiger aus den Fenster läuft, entferne das Element
		for que.Len() > 0 && que.Front().Value.(int) <= i-k {
			que.Remove(que.Front())
		}

		// Entferne das letzte Elemente, wenn es kleiner als das aktuelle Element ist
		for que.Len() > 0 && v[que.Back().Value.(int)] <= n {
			que.Remove(que.Back())
		}

		// Füge das aktuelle Element hinzu
		que.PushBack(i)

		// Füge das Maximum zum Ergebnis hinzu
		if i >= k-1 {
			result = append(result, v[que.Front().Value.(int)])
		}
	}
	end := time.Now()

	//Ausgabe
	fmt.Println("V = ", v)
	fmt.Println("k = ", k)
	fmt.Println(" => ", result)
	fmt.Println("needed: ", end.Sub(start))
}
