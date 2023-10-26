package main

import (
	"errors"
	"fmt"
)

var ErrNoCoords = errors.New("no coordinates given")
var ErrCoordsOutOfRange = errors.New("coordinates out of range")
var ErrBoardToSmall = errors.New("n must be greater than 0")
var ErrNotEnoughCoords = errors.New("not enough coordinates")
var ErrToMuchCoords = errors.New("to much coordinates")

type Coord struct {
	i, j int
}

func main() {
	var c []Coord

	n := 3
	c = []Coord{{0, 0}, {0, 2}, {1, 1}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {0, 1}, {0, 2}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {1, 0}, {2, 0}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{2, 0}, {0, 2}, {1, 1}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{0, 2}, {1, 1}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{}
	CheckGame(n, c)
	c = []Coord{{1, 1}, {2, 2}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {0, 2}, {1, 1}, {2, 3}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {0, 2}, {1, 1}, {2, -2}}
	CheckGame(n, c)
	c = []Coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}, {0, 0}}
	CheckGame(n, c)
}

// Baut das Spiel auf und überprüft ob jemand gewonnen hat
func CheckGame(n int, c []Coord) (bool, error) {
	err := checkInput(n, c)
	if err != nil {
		fmt.Println("error: ", err)
		return false, err
	}

	b := createBoard(n)
	fillBoard(b, c)
	res := checkWin(b)

	printBoard(b)
	if res {
		fmt.Println("Player won!")
	} else {
		fmt.Println("No winner!")
	}
	fmt.Println()

	return res, nil
}

// Überprüft ob die Eingaben korrekt sind
func checkInput(n int, c []Coord) error {
	if len(c) == 0 {
		return ErrNoCoords
	}

	if len(c) < n {
		return ErrNotEnoughCoords
	}

	if len(c) > n*n {
		return ErrToMuchCoords
	}

	if n <= 0 {
		return ErrBoardToSmall
	}

	for _, v := range c {
		if v.i < 0 || v.i >= n || v.j < 0 || v.j >= n {
			return ErrCoordsOutOfRange
		}
	}
	return nil
}

// Erstellt ein 2D Array mit der Größe n*n und füllt es mit 0
func createBoard(n int) [][]int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	return board
}

// Füllt das Board mit den Koordinaten
func fillBoard(b [][]int, c []Coord) [][]int {
	for _, v := range c {
		b[v.i][v.j] = 1
	}
	return b
}

// Gibt das Board aus
func printBoard(b [][]int) {
	//Header (X-Achse)
	fmt.Print("  ")
	for i := 0; i < len(b); i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//Ausgabe der Koordinaten
	for i := 0; i < len(b); i++ {
		fmt.Print(i, " ")
		for j := 0; j < len(b); j++ {
			if b[i][j] == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

// Überprüft ob jemand gewonnen hat
func checkWin(b [][]int) bool {
	var d1, d2 int

	for i := 0; i < len(b); i++ {
		var h, v int

		for j := 0; j < len(b); j++ {
			//Horizontal
			h += b[i][j]
			//Vertikal
			v += b[j][i]
			//Diagonal \
			if i == j {
				d1 += b[i][j]
			}
			//Diagonal /
			if i+j == len(b)-1 {
				d2 += b[i][j]
			}
		}

		if h == len(b) || v == len(b) || d1 == len(b) || d2 == len(b) {
			return true
		}
	}

	return false
}
