package main

import "fmt"

type Cord struct {
	i, j int
}

func main() {
	n := 3
	c := []Cord{{0, 0}, {0, 2}, {1, 1}, {2, 2}}

	board := createBoard(n)
	fillBoard(board, c)

	printBoard(board)
}

func createBoard(n int) [][]int {
	//Erstellt ein 2D Array mit der Größe n*n und füllt es mit 0
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	return board
}

func fillBoard(b [][]int, c []Cord) [][]int {
	//Füllt das Board mit den Koordinaten
	for _, v := range c {
		b[v.i][v.j] = 1
	}
	return b
}

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
