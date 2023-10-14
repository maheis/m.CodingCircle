package main

import (
	"fmt"

	"github.com/maheis/CodingCircle/231011_Median/bstree"
	"github.com/maheis/CodingCircle/231011_Median/chained"
	"github.com/maheis/CodingCircle/231011_Median/heap"
	"github.com/maheis/CodingCircle/231011_Median/slice"
)

// 11.10.2023 - https://www.youtube.com/watch?v=28ad9BYDHlI

//Aufgabenstellung:
// Gesucht ist eine Funktion, die für eine Liste von Zahlen den
// kontinuierlichen Median berechnet, und diese Werte als Liste zurückgibt.
//
//Beispiel:
// [17, 2, 8, 27, 12, 9]
// => [17, 9.5, 8, 12.5, 12, 10.5]
//
//Definition Median
// Der Wert, der genau in der Mitte einer Datenverteilung liegt, nennt sich Median oder Zentralwert.
// Die eine Hälfte aller Individualdaten ist immer kleiner, die andere größer als der Median.
// Bei einer geraden Anzahl von Individualdaten ist der Median die Hälfte der Summe der beiden in der Mitte liegenden Werte.
// Beispiele:
//  1,2,3,5,6,7,12: Der Median ist 5.
//  1,1,1,1,12: Der Median ist 1.
//  1,2,3,4,5,9: Der Median ist 3,5 (3 und 4 liegen um die Mitte – die Hälfte der Summe 7 ist 3,5).
// Quelle: https://de.statista.com/statistik/lexikon/definition/85/median/#:~:text=Der%20Wert%2C%20der%20genau%20in,in%20der%20Mitte%20liegenden%20Werte.
//

func main() {
	//Aufgabestellung:
	values := []int{17, 2, 8, 27, 12, 9}
	// values := []int{17, 2, 8, 27, 12, 9, 13, 25, 29, 1, 19, 20, 33, 42, 6, 11, 49, 18}

	//Random
	// cnt := 33333
	// values := make([]int, 0, cnt)
	// for i := 0; i < cnt; i++ {
	// 	values = append(values, rand.Intn(100))
	// }

	if len(values) < 42 {
		fmt.Println("values: ", values)
	}

	slice.Run(values)
	chained.Run(values)
	bstree.Run(values)
	heap.Run(values)
}
