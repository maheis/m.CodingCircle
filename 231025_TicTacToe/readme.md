# 25.10.2023 - TicTacToe
https://www.youtube.com/watch?v=ofs5v4jXa-k

# Aufgabenstellung:
Gegeben seien die Spielfeldlänge n und eine Liste von Koordinaten.
Gesucht ist eine Funktion, die für ein Spielfeld der Größe n*n ent-
scheidet, ob die Koordinaten eine Gewinnposition enthalten.

## Beispiel:
```
n = 3; 
coords [
    (0,0), (0,2), (1,1), (2,2)
] => true
```

## Bedingungen:
Laufzeit O(n²), Speicher O(n), Single-Pass
<sub>Golo: Singlepass heißt, dass Du eigentlich nur eine einzige Schleife haben solltest, in der Du jedes Element einmal triffst. Da das Spielfeld aber quadratisch ist, entspricht die Anzahl der Schleifendurchläufe – je nach Implementierung – eventuell nicht n, sondern n^2. Daher die Forderung mit quadratischer Laufzeit.</sub>

# Skizze:
```
Beispiel:
    0 1 2
  0 X   X 
  1   X  
  2     X

Mögliche Lösungskoordinaten:
  Horizontal:
    0,0, 0,1 0,2
    1,0, 1,1 1,2
    2,0, 2,1 2,2
  Diagonal:
    0,0, 1,1 2,2
    0,2, 1,1 2,0
  Vertikal:
    0,0, 1,0 2,0
    0,1, 1,1 2,1
    0,2, 1,2 2,2
```

# Lösung:
## Board
Spielfeld aufbauen und in 3 Wegen durchlaufen:
- Horizontal
- Vertikal
- Diagonal

## Koordinaten-Summe
Zählen wie oft eine Spalte, Zeile oder Diagonale besetzt ist. Wenn die Anzahl der Besetzungen gleich der Spielfeldgröße ist, dann ist die Gewinnbedingung erfüllt. 
Zugrunde liegt der Lösung die Gewinnregeln von TicTacToe. Es gibt 3 Gewinnmöglichkeiten:
- Horizontal linie
- Vertikal linie
- Diagonal linie
Setzt aber vorraus das keien Koordinate doppelt vorkommt!

# Notizen:
Kann man nicht Binärsummen verwenden? Die sind ja Eindeutig und könnten im laufen berechnet werden.
```
n = 3
      0   1   2
  0   1   2   4 =   7
  1   8  16  32 =  56
  2  64 128 256 = 448
  =   =   =   = =
84   73 146 292   273
```

# Performance:

## Szenario 1
Beide Routinen müssen nur eine Zeile durchlaufen.
In dem Fall würde ich erwarten das die Performance beider Routinen nahezu gleich ist.
```
  0 1 2
0 X X X
1  
2 
```

## Szenario 2
Die eine Routine muss eine Zeile durchlaufen, die andere Routine muss alle Felder durchlaufen!
In dem Fall würde ich erwarten das die Performance stark differiert.
```
  0 1 2
0
1
2 X X X
```

## Szenario 3
Die eine Routine muss einen Coordinatensatz durchlaufen, die andere Routine muss alle Felder durchlaufen!
In dem Fall würde ich erwarten das die Performance stark differiert.
```
  0 1 2
0 X 
1   X
2     X
```