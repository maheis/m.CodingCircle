package bstree

import (
	"fmt"
	"time"
)

type bstNode struct {
	val   int
	count int
	left  *bstNode
	right *bstNode
}

func Run(values []int) {
	fmt.Println("bstree")
	start := time.Now()
	medians := ContinuousMedian(values)
	end := time.Now()
	if len(medians) < 42 {
		fmt.Println(" => medians: ", medians)
	}
	fmt.Println(" => needed:", end.Sub(start))
}

func initBST() *bstNode {
	root := &bstNode{val: 9223372036854775807, count: 0, left: nil, right: nil}

	return root
}

func appendBSTree(root *bstNode, val int) *bstNode {
	// Wenn der Baum leer ist, dann erzeuge einen neuen Knoten
	if root == nil {
		return &bstNode{val: val}
	}

	// Ist der Wert kleiner als der Wert des aktuellen Knotens, dann gehe nach links
	if val < root.val {
		root.left = appendBSTree(root.left, val)
	} else {
		// Ist der Wert größer als der Wert des aktuellen Knotens, dann gehe nach rechts
		root.right = appendBSTree(root.right, val)
	}

	return root
}

func printBSTree(root *bstNode) {
	// Wenn der Knoten leer ist, dann beende die Funktion
	if root == nil {
		return
	}

	// Gehe rekursiv nach links bis der linke Knoten leer ist
	printBSTree(root.left)
	// Gib den Wert des linken Knotens aus
	fmt.Print(root.val, " -> ")
	// Gehe rekursiv eins nach rechts bis der rechte Knoten irgendwann auch leer ist
	printBSTree(root.right)

	// Auf diese Weise wird der Baum in aufsteigender Reihenfolge ausgegeben. Alles was kleiner ist als der aktuelle Knoten, steht links, sprich alles kleiner dem aller ersten Knoten Links, sonst alles immer nach rechts geschubst!
}

func searchBSMedian(root *bstNode) (int, int) {
	count := root.count / 2
	valA := -1
	valB := -1

	findBSMedian(root, &count, &valA, &valB)

	return valA, valB
}

func findBSMedian(root *bstNode, count *int, valA *int, valB *int) {
	// Wenn der Knoten leer ist, dann beende die rekursion
	if root == nil {
		return
	}

	findBSMedian(root.left, count, valA, valB)
	*count = *count - 1
	if *count < -1 {
		return
	}
	if *count == -1 {
		*valB = root.val
		return
	}
	if *count == 0 {
		*valA = root.val
	}
	findBSMedian(root.right, count, valA, valB)
}

func ContinuousMedian(values []int) []float64 {
	// Partial das Array aufspalten um den Kontinuierlichen Median zu berechnen
	partitions := initBST()
	// Liste der Medians
	medians := make([]float64, 0, len(values))

	for _, value := range values {
		partitions = appendBSTree(partitions, value)
		partitions.count++
		medians = append(medians, Median(partitions))
	}

	return medians
}

func Median(root *bstNode) float64 {
	count := root.count //Anzahl der Elemente in der Liste

	//2 Elemente aus der Liste abholen...
	valA, valB := searchBSMedian(root)

	// Grade
	if count%2 == 0 {
		// Mittelwert der beiden mittleren Werte

		//HIER muss der Vorgänger, nicht der Nachfolger genommen werden.
		return float64(valA+valB) / 2
	}

	// Ungrade
	return float64(valB)
}
