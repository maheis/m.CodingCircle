package main

import "fmt"

// 18.10.2023 - https://www.youtube.com/watch?v=0QPK9X1PtEE

//Aufgabenstellung:
// Gegeben sind eine Liste V und eine Ganzzahl k. Ermittle das laufende maximum für die Fenstergröße k.
//
//Beispiel:
// V = [ 27, 9, 17, 2, 12, 8 ]
// k = 3
// => [ 27, 17, 17, 12 ]
//
//Bedingungen:
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
//

//

func main() {
	values := []int{27, 9, 17, 2, 12, 8}
	fmt.Println("V = ", values)
	k := 3
	fmt.Println("k = ", k)

	var partitions []int
	var maxis []int

	if len(values) > k {
		// Anfang des ersten Fenster erstellen
		for i := 0; i < k-1; i++ {
			partitions = append(partitions, values[i])
		}

		for i := k - 1; i < len(values); i++ {
			// Wert ins Fenster einfügen
			partitions = append(partitions, values[i])

			// Maximum im Fenster finden
			max := 0
			for _, v := range partitions {
				if v > max {
					max = v
				}
			}
			maxis = append(maxis, max)

			// Ein Element aus dem Fenster entfernen
			partitions = partitions[1:]

		}
	}

	fmt.Println(" => ", maxis)
}
