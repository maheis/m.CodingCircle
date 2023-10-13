package main

//Mögliche Lösung über eine Map
// Über eine Map nicht sinnvoll zu realisieren. Map müsste für Sort in ein Slice überführt werden...
// Würde man das machen, wäre es aber auch nicht schneller als die slice-Variante.
// Es könnte der Key als Value verwendet werden. Dann würde sich die Map automatisch sortieren.
// Die Map müsste aber von vorne durchlaufen werden, außerdem gehen keine doppelten Werte.

//Mögliche Lösung über zwei Listen
// Die Idee ist, dass wir die Liste der Zahlen in zwei Hälften aufteilen.
// Die linke Hälfte enthält alle Zahlen, die kleiner als der Median sind.
// Die rechte Hälfte enthält alle Zahlen, die größer als der Median sind.
// Die linke Hälfte wird als Liste implementiert, die rechte Hälfte als Liste.
// Der Median ist dann entweder das letzte Element der linken Liste oder das erste Element der rechten Liste.
// Wenn die Anzahl der Zahlen gerade ist, dann ist der Median der Durchschnitt der beiden Elemente.

//Sortierter Baum ?
//

//Idee: Kette mit Forward/Backward spulen zu können?!

//Grundsätzlcihe Frage, braucht man alle Werte um den Median zu berechnen?
//Oder Reicht ein Subset um den Median herum aus?
