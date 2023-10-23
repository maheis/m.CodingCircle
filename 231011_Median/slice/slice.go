package slice

import (
	"fmt"
	"sort"
	"time"
)

func Run(values []int) {
	fmt.Println("slice")
	start := time.Now()
	medians := ContinuousMedian(values)
	end := time.Now()
	if len(medians) < 42 {
		fmt.Println(" => medians: ", medians)
	}
	fmt.Println(" => needed:", end.Sub(start))
}

func ContinuousMedian(values []int) []float64 {
	// Partial das Array aufspalten um den Kontinuierlichen Median zu berechnen
	partitions := make([]int, 0, len(values))
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	for _, value := range values {
		partitions = append(partitions, value)
		sort.Ints(partitions) //Sotierte Liste für den Median
		medians = append(medians, Median(partitions))
	}

	return medians
}

func Median(values []int) float64 {
	n := len(values)

	// Grade
	if n%2 == 0 {
		// Mittelwert der beiden mittleren Werte
		return float64(values[n/2-1]+values[n/2]) / 2.0
	}
	// Ungrade
	return float64(values[n/2])
}
