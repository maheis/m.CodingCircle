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

func main() {
	v := []int{27, 9, 17, 2, 12, 8}
	fmt.Println("V = ", v)
	k := 3
	fmt.Println("k = ", k)

	var maxis []int

	if len(v) > k {
		// Schleife soll abzüglich der Fenstergröße k laufen
		for i := 0; i < len(v)-k+1; i++ {
			// Maximalwert im Fensters ermitteln, dazu die Fenstergröße voreilen
			max := v[i]
			for j := i + 1; j < i+k; j++ {
				if v[j] > max {
					max = v[j]
				}
			}
			maxis = append(maxis, max)

			// Im nächsten durchlauf wird das Fenster 1 nach vorne geschoben...
		}
	}

	fmt.Println(" => ", maxis)
}
