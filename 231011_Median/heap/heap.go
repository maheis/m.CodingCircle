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

func NewMaxHeap() MaxHeap {
	h := MaxHeap{}
	heap.Init(&h)
	return h
}
func (h MaxHeap) Len() int           { return len(h) }            // Länge des Heaps
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }       // Vergleichsfunktion von zwei Elementen; Elemente wander nach oben im MaxHeap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }  // Tausche zwei Elemente; spezielle Schreibweise in Go: h[i] wird zu h[j] und h[j] zu h[i], ohne dass eine temporäre Variable benötigt wird
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) } // Neues Element hinzufügen; spezielle Schreibweise in Go: x.(int) ist ein Typecast
func (h *MaxHeap) Pop() any {
	old := *h         // Kopie des Heaps bzw. Pointer in eine Variable überführen (dereferenzieren)
	n := len(old) - 1 // Index auf das letzte Element
	*h = old[0:n]     // Alles bis auf das letzte Element als Rückgabe in den (ByRef) Heap zurückschreiben
	return old[n]     // Letztes Element zurückgeben
}

type MinHeap []int

func NewMinHeap() MinHeap {
	h := MinHeap{}
	heap.Init(&h)
	return h
}
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Unterschied zum MaxHeap! Elemente wander nach unten im MinHeap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	*h = old[0:n]
	return old[n]
}

func ContinuousMedian(values []int) []float64 {
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	maxHeap := NewMaxHeap() // linke Hälfte, sortiert absteigend (höchstes Element ist das erste Element)
	minHeap := NewMinHeap() // rechte Hölfte, sortiert aufsteigend (niedrigstes Element ist das erste Element)

	for _, num := range values {
		// Value in den Heap packen, der weniger Elemente hat
		if maxHeap.Len() <= minHeap.Len() {
			heap.Push(&maxHeap, num)
		} else {
			heap.Push(&minHeap, num)
		}

		// Balanciere die Heaps aus
		if maxHeap.Len()-minHeap.Len() > 1 {
			heap.Push(&minHeap, heap.Pop(&maxHeap))
		} else if minHeap.Len()-maxHeap.Len() > 1 {
			heap.Push(&maxHeap, heap.Pop(&minHeap))
		}

		// Ermittel den Median
		if maxHeap.Len() == minHeap.Len() { // gerade Anzahl an Elementen
			medians = append(medians, float64(maxHeap[0]+minHeap[0])/2.0)
		} else if maxHeap.Len() > minHeap.Len() { // ungerade Anzahl an Elementen
			medians = append(medians, float64(maxHeap[0]))
		} else { // ungerade Anzahl an Elementen
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
