package main

//Mögliche Lösung über eine Map
// Über eine Map nicht sinnvoll zu realisieren. Map müsste für Sort in ein Slice überführt werden...
// Würde man das machen, wäre es aber auch nicht schneller als die slice-Variante.
// Es könnte der Key als Value verwendet werden. Dann würde sich die Map automatisch sortieren.
// Die Map müsste aber von vorne durchlaufen werden, außerdem gehen keine doppelten Werte.

// AVL-Baum ?

//Idee: Kette mit Forward/Backward spulen zu können?!

//Grundsätzlcihe Frage, braucht man alle Werte um den Median zu berechnen?
//Oder Reicht ein Subset um den Median herum aus?
