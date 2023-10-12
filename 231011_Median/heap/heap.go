package heap

// 11.10.2023 - https://www.youtube.com/watch?v=28ad9BYDHlI

//Aufgabenstellung:
// Gesucht ist eine Funktion, die für eine Liste von Zahlen den
// kontinuierlichen Median berechnet, und diese Werte als Liste zurückgibt.
//
//Beispiel:
// [17, 2, 8, 27, 12, 9]
// => [17, 9.5, 8, 12.5, 12, 10.5]
//
//Lösung:
// Die Idee ist, dass wir die Liste der Zahlen in zwei Hälften aufteilen.
// Die linke Hälfte enthält alle Zahlen, die kleiner als der Median sind.
// Die rechte Hälfte enthält alle Zahlen, die größer als der Median sind.
// Die linke Hälfte wird als Max-Heap implementiert, die rechte Hälfte als Min-Heap.
// Der Median ist dann entweder die Wurzel des Max-Heaps oder die Wurzel des Min-Heaps.
// Wenn die Anzahl der Zahlen gerade ist, dann ist der Median der Durchschnitt der beiden Wurzeln.
//
