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
	root := &bstNode{val: 0, count: 0, left: nil, right: nil}

	return root
}

func printBSTree(root *bstNode) {
	// Wenn der Knoten leer ist, dann beende die Funktion
	if root == nil {
		return
	}

	// Gehe rekursiv nach links bis der Knoten leer ist
	printBSTree(root.left)
	// Gib den Wert des Knotens aus
	fmt.Print(root.val, " -> ")
	// Gehe rekursiv nach rechts bis der Knoten leer ist
	printBSTree(root.right)

	// Auf diese Weise wird der Baum in aufsteigender Reihenfolge ausgegeben. Alles was kleiner ist als der aktuelle Knoten, steht links, sprich alles kleiner dem aller ersten Knoten Links, sonst alles immer nach rechts geschubst!
}

func nextBSNode(root *bstNode) *bstNode {
	// Gehe rekursiv nach links bis der Knoten leer ist
	root = root.left
	if root == nil {
		// Gehe rekursiv nach rechts bis der Knoten leer ist
		root = root.right
	}

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

func ContinuousMedian(vals []int) []float64 {
	// Partial das Array aufspalten um den Kontinuierlichen Median zu berechnen
	partitions := initBST()
	// Liste der Medians
	medians := make([]float64, 0, len(vals))

	for i := 0; i < len(vals); i++ {
		partitions = appendBSTree(partitions, vals[i])
		partitions.count++
		medians = append(medians, Median(partitions))
	}

	return medians
}

func Median(root *bstNode) float64 {
	count := root.val //Anzahl der Elemente in der Liste

	//in die Mitte der Liste gehen
	for i := 0; i <= count/2; i++ {
		root = nextBSNode(root)
	}

	// Grade
	if count%2 == 0 { //Head.val enthält den Count der Liste!
		// Mittelwert der beiden mittleren Werte
		return float64(root.val+nextBSNode(root).val) / 2
	}

	// Ungrade
	return float64(root.val)
}
