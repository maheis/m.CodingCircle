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
<sub>Golo: Singlepass heißt, dass Du eigentlich nur eine einzige Schleife haben solltest, in der Du jedes Element einmal triffst. Da das Spielfeld aber quadratisch ist, entspricht die Anzahl der Schleifendurchläufe – je nach Implementierung – eventuell nicht n, sondern n^2. Daher die Forderung mit quadratischer Laufzeit.</sub>

# Skizze:
```
Beispiel:
    0 1 2
  0 #   # 
  1   #  
  2     #

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
Spielfeld aufbauen und in 3 Wegen durchlaufen:
- Horizontal
- Vertikal
- Diagonal

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