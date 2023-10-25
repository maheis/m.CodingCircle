package main

import (
	"container/list"
	"errors"
	"fmt"
	"time"
)

// 18.10.2023 - https://www.youtube.com/watch?v=0QPK9X1PtEE

// Aufgabenstellung:
// Gegeben sind eine Liste V und eine Ganzzahl k. Ermittle das laufende maximum für die Fenstergröße k.
//
// Beispiel:
// V = [ 27, 9, 17, 2, 12, 8 ]
// k = 3
// => [ 27, 17, 17, 12 ]
//
// Bedingungen:
// Die Laufzeit soll O(n) betragen, der Speicherbedarf soll O(k) betragen.
//
// Hinweis:
// Ein Algorithmus mit einer Laufzeit von O(n) hat eine lineare Laufzeit, die proportional zur Größe
// der Eingabe ist. Das bedeutet, dass die Laufzeit des Algorithmus linear mit der Größe der Eingabe wächst.
// Wenn die Größe der Eingabe um den Faktor k erhöht wird, sollte die Laufzeit des Algorithmus auch um den
// Faktor k erhöht werden.
// Der Speicherbedarf von O(k) bedeutet, dass der Speicherbedarf des Algorithmus proportional zur Größe
// des Eingabeparameters k ist. Mit anderen Worten, wenn die Größe des Eingabeparameters um den Faktor k
// erhöht wird, sollte der Speicherbedarf des Algorithmus auch um den Faktor k erhöht werden.
//
// Lösung:
// Es wird das Fenster betrachtet nicht die Werte. Dazu wird eine Que verwendet, die die Indizes der
// Elemente enthält. Die Que ist sortiert, so dass das erste Element immer das größte Element des aktuellen Fensters ist.
// Das nachfolgende Element enthält das größte Element für das Nachfolgende Fenster.
// Wenn das Fenster sich bewegt, wird das erste Element entfernt, wenn es aus dem Fenster läuft.
// Wenn ein neues Element hinzugefügt wird, wird das letzte Element entfernt, wenn es kleiner als das neue Element ist.
// Die Liste enthält maximal k Elemente, so dass der Speicherbedarf O(k) ist.
// Die Liste wird maximal n mal durchlaufen, so dass die Laufzeit O(n) ist.
//

func main() {
	v := []int{27, 9, 17, 2, 12, 8}
	// 		   27, 9, 17
	// 			   9, 17, 2
	//				  17, 2, 12
	// 					  2, 12, 8
	// => [ 27 17 17 12 ]
	k := 3

	// v := []int{27, 9, 17, 2, 12, 8, 9, 10, 25, 29, 3, 8, 9, 8, 10}
	// => [27 17 17 12 12 10 25 29 29 29 9 9 10]
	// k := 3

	//Random
	// cnt := 3333333
	// v := make([]int, 0, cnt)
	// for i := 0; i < cnt; i++ {
	// 	v = append(v, rand.Intn(100))
	// }
	// k := 3333

	if len(v) < 42 {
		fmt.Println("V = ", v)
	}
	fmt.Println("k = ", k)

	iterative(v, k)
	que(v, k)
	slice(v, k)
}

func iterative(v []int, k int) ([]int, error) {
	if len(v) == 0 {
		return []int{}, errors.New("v is empty")
	}
	if k <= 0 {
		return []int{}, errors.New("k must be greater than 0")
	}
	if k > len(v) {
		return []int{}, errors.New("k must be less than len(v)")
	}

	result := make([]int, 0, len(v)-k+1)

	start := time.Now()
	// Schleife soll abzüglich der Fenstergröße k laufen
	for i := 0; i < len(v)-k+1; i++ {
		// Maximalwert im Fensters ermitteln, dazu die Fenstergröße voreilen
		max := v[i]
		for j := i + 1; j < i+k; j++ {
			if v[j] > max {
				max = v[j]
			}
		}
		result = append(result, max)

		// Im nächsten durchlauf wird das Fenster 1 nach vorne geschoben...
	}
	end := time.Now()

	//Ausgabe
	fmt.Println("iterative")
	if len(result) < 42 {
		fmt.Println(" => ", result)
	}
	fmt.Println("needed: ", end.Sub(start))

	return result, nil
}

func que(v []int, k int) ([]int, error) {
	if len(v) == 0 {
		return []int{}, errors.New("v is empty")
	}
	if k <= 0 {
		return []int{}, errors.New("k must be greater than 0")
	}
	if k > len(v) {
		return []int{}, errors.New("k must be less than len(v)")
	}

	result := make([]int, 0, len(v)-k+1)

	var que list.List

	start := time.Now()
	//Itteriere durch die Liste...
	for i, n := range v {
		// Wenn der Index aus den Fenster läuft, entferne diesen
		if que.Len() > 0 && que.Front().Value.(int) <= i-k {
			que.Remove(que.Front())
		}

		// Entferne den letzten Index, wenn der Wert auf den er Zeigt kleiner als das aktuelle Wert ist
		for que.Len() > 0 && v[que.Back().Value.(int)] <= n {
			que.Remove(que.Back())
		}

		// Füge den Index des aktuelle Element hinzu
		que.PushBack(i)

		// Sobald das erste Fenster abgeschlossen ist füge das Maximum zum Ergebnis hinzu
		if i >= k-1 {
			result = append(result, v[que.Front().Value.(int)])
		}
	}
	end := time.Now()

	//Ausgabe
	fmt.Println("que")
	if len(result) < 42 {
		fmt.Println(" => ", result)
	}
	fmt.Println("needed: ", end.Sub(start))

	return result, nil
}

func slice(v []int, k int) ([]int, error) {
	if len(v) == 0 {
		return []int{}, errors.New("v is empty")
	}
	if k <= 0 {
		return []int{}, errors.New("k must be greater than 0")
	}
	if k > len(v) {
		return []int{}, errors.New("k must be less than len(v)")
	}

	result := make([]int, 0, len(v)-k+1)

	var que []int

	start := time.Now()
	//Itteriere durch die Liste...
	for i := 0; i < len(v); i++ {
		// Wenn der Index aus den Fenster läuft, entferne diesen
		if len(que) > 0 && que[0] <= i-k {
			que = que[1:] // Slice kopieren ohne das 0te Element (Elemente 1 bis len(que))
		}

		// Entferne den letzten Index, wenn der Wert auf den er Zeigt kleiner als das aktuelle Wert ist
		for len(que) > 0 && v[que[len(que)-1]] <= v[i] {
			que = que[:len(que)-1] // Slice kopieren ohne das letzte Element (Elemente 0 bis len(que)-1)
		}

		// Füge den Index des aktuelle Element hinzu
		que = append(que, i)

		// Sobald das erste Fenster abgeschlossen ist füge das Maximum zum Ergebnis hinzu
		if i >= k-1 {
			result = append(result, v[que[0]])
		}
	}
	end := time.Now()

	//Ausgabe
	fmt.Println("slice")
	if len(result) < 42 {
		fmt.Println(" => ", result)
	}
	fmt.Println("needed: ", end.Sub(start))

	return result, nil
}
