package main

import (
	"fmt"
	"sort"
)

// 11.10.2023 - https://www.youtube.com/watch?v=28ad9BYDHlI

//Aufgabenstellung:
// Gesucht ist eine Funktion, die für eine Liste von Zahlen den
// kontinuierlichen Median berechnet, und diese Werte als Liste zurückgibt.
//
//Beispiel:
// [17, 2, 8, 27, 12, 9]
// => [17, 9.5, 8, 12, 12, 10.5]
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
	values := []int{17, 2, 8, 27, 12, 9}
	fmt.Println("values: ", values)

	medians := continuousMedian(values)
	fmt.Println(" => medians: ", medians)
}

func continuousMedian(values []int) []float64 {
	// Partial das Array aufspalten um den Kontinuierlichen Median zu berechnen
	partitions := make([]int, 0, len(values))
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	for i := 0; i < len(values); i++ {
		partitions = append(partitions, values[i])
		sort.Ints(partitions) //Sotierte Liste für den Median
		medians = append(medians, getMedian(partitions))
	}

	return medians
}

func getMedian(values []int) float64 {
	// Grade
	if len(values)%2 == 0 {
		// Mittelwert der beiden mittleren Werte
		return float64(values[len(values)/2-1]+values[len(values)/2]) / 2
	}
	// Ungrade
	return float64(values[len(values)/2])
}
