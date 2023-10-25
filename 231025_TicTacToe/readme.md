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
Laufteit O(n²), Speicher O(n), Single-Pass

# Notizen
```
Beispiel:
    0 1 2
  0 X    
  1   X  
  2 X   X

Mögliche Lösungskoordinaten:
  Vertikal:
    0,0, 0,1 0,2
    1,0, 1,1 1,2
    2,0, 2,1 2,2
  Diagonal:
    0,0, 1,1 2,2
    0,2, 1,1 2,0
  Horizontal:
    0,0, 1,0 2,0
    0,1, 1,1 2,1
    0,2, 1,2 2,2

n = 5
  0 1 2 3 4 5
  1 - - - - -
  2 - - - - -
  3 - - - - -
  4 - - - - -
  5 - - - - -
 ```

# Lösung