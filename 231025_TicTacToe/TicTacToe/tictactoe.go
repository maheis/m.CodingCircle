package tictactoe

import (
	"errors"
	"fmt"
)

var ErrNoCoords = errors.New("no coordinates given")
var ErrBoardToSmall = errors.New("n must be greater than 0")
var ErrNotEnoughCoords = errors.New("not enough coordinates")
var ErrToMuchCoords = errors.New("to much coordinates")
var ErrCoordsOutOfRange = errors.New("coordinates out of range")

type Coord struct {
	X, Y int
}

// CheckGame Baut das Spiel auf und überprüft ob jemand gewonnen hat
func CheckGame(n int, c []Coord) (bool, error) {
	err := checkInput(n, c)
	if err != nil {
		fmt.Println("error: ", err)
		return false, err
	}

	b := createBoard(n)
	fillBoard(b, c)
	res := checkWinBoard(b)

	printBoard(b)
	if res {
		fmt.Println("Player won!")
	} else {
		fmt.Println("No winner!")
	}
	fmt.Println()

	return res, nil
}

// checkInput überprüft ob die Eingaben korrekt sind
func checkInput(n int, c []Coord) error {
	if len(c) == 0 {
		return ErrNoCoords
	}

	if n <= 0 {
		return ErrBoardToSmall
	}

	if len(c) < n {
		return ErrNotEnoughCoords
	}

	if len(c) > n*n {
		return ErrToMuchCoords
	}

	for _, v := range c {
		if v.X < 0 || v.X >= n || v.Y < 0 || v.Y >= n {
			return ErrCoordsOutOfRange
		}
	}
	return nil
}

// createBoard erstellt ein 2D Array mit der Größe n*n und füllt es mit 0
func createBoard(n int) [][]int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	return board
}

// fillBoard füllt das Board mit den Koordinaten
func fillBoard(b [][]int, c []Coord) [][]int {
	for _, v := range c {
		b[v.X][v.Y] = 1
	}
	return b
}

// printBoard gibt das Board aus
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

// checkWinBoard überprüft ein Übergebenes Board auf den Gewinn
func checkWinBoard(b [][]int) bool {
	var d1, d2 int // Diagonal \, Diagonal /

	for i := 0; i < len(b); i++ {
		var h, v int //Horizontal, Vertikal

		for j := 0; j < len(b); j++ {
			h += b[i][j]
			v += b[j][i]
			if i == j {
				d1 += b[i][j]
			}
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

// checkWinCoords überprüft ein Übergebenes Board n und die Eingabekoodrinaten auf den Gewinn
func checkWinCoords(n int, c []Coord) bool {
	h := make([]int, n) //Horizontal
	v := make([]int, n) //Vertikal
	d1 := 0             // Diagonal \
	d2 := 0             // Diagonal /

	for _, field := range c {
		h[field.X]++
		v[field.Y]++
		if field.X == field.Y {
			d1++
		}
		if field.X+field.Y == n-1 {
			d2++
		}

		if h[field.X] == n || v[field.Y] == n || d1 == n || d2 == n {
			return true
		}
	}
	return false
}
