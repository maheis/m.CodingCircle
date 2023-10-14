package heap

// Die Idee ist, dass wir die Liste der Zahlen in zwei Hälften aufteilen.
// Die linke Hälfte enthält alle Zahlen, die kleiner als der Median sind.
// Die rechte Hälfte enthält alle Zahlen, die größer als der Median sind.
// Die linke Hälfte wird als heap implementiert, die rechte Hälfte als heap.
// Der Median ist dann entweder das letzte Element des linken hep oder das erste Element des rechten heap.
// Wenn die Anzahl der Zahlen gerade ist, dann ist der Median der Durchschnitt der beiden Elemente.

import (
	"container/heap"
	"fmt"
	"time"
)

type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ContinuousMedian(values []int) []float64 {
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	var maxHeap MaxHeap
	var minHeap MinHeap
	heap.Init(&maxHeap)
	heap.Init(&minHeap)

	for _, num := range values {
		if maxHeap.Len() == 0 || num <= maxHeap[0] {
			heap.Push(&maxHeap, num)
		} else {
			heap.Push(&minHeap, num)
		}

		if maxHeap.Len()-minHeap.Len() > 1 {
			heap.Push(&minHeap, heap.Pop(&maxHeap))
		} else if minHeap.Len()-maxHeap.Len() > 1 {
			heap.Push(&maxHeap, heap.Pop(&minHeap))
		}

		if maxHeap.Len() == minHeap.Len() {
			medians = append(medians, float64(maxHeap[0]+minHeap[0])/2.0)
		} else if maxHeap.Len() > minHeap.Len() {
			medians = append(medians, float64(maxHeap[0]))
		} else {
			medians = append(medians, float64(minHeap[0]))
		}
	}

	return medians
}

func Run(values []int) {
	fmt.Println("heap")
	start := time.Now()
	medians := ContinuousMedian(values)
	end := time.Now()
	if len(medians) < 42 {
		fmt.Println(" => medians: ", medians)
	}
	fmt.Println(" => needed:", end.Sub(start))
}
