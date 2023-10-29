package tictactoe

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func Performance() {
	print := false
	rounds := 99

	// ## Szenario 1
	// Beide Routinen müssen nur eine Zeile durchlaufen.
	// In dem Fall würde ich erwarten das die Performance beider Routinen nahezu gleich ist.
	//   0 1 2
	// 0 X X X
	// 1
	// 2
	performance(szenario1, print, rounds)

	// ## Szenario 2
	// Die eine Routine muss eine Zeile durchlaufen, die andere Routine muss alle Felder durchlaufen!
	// In dem Fall würde ich erwarten das die Performance stark differiert.
	//   0 1 2
	// 0
	// 1
	// 2 X X X
	performance(szenario2, print, rounds)

	// ## Szenario 3
	// Die eine Routine muss einen Coordinatensatz durchlaufen, die andere Routine muss alle Felder durchlaufen!
	// In dem Fall würde ich erwarten das die Performance stark differiert.
	//   0 1 2
	// 0 X
	// 1   X
	// 2     X
	performance(szenario3, print, rounds)
}

func performance(f func(int) []Coord, print bool, rounds int) {
	var n int
	var c []Coord
	var b [][]int

	fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())

	if print {
		n = 3
		c = f(n)
		b = createBoard(n)
		fillBoard(b, c)
		printBoard(b)
		checkWinBoard(b)
		checkWinCoords(n, c)
	}

	n = 0
	for i := 0; i < rounds; i++ {
		n = n + 100
		c = f(n)
		b = createBoard(n)
		stopwatch(n, c, b)
	}

	fmt.Println("")
}

func szenario1(n int) []Coord {
	c := make([]Coord, n)

	for i := 0; i < n; i++ {
		c = append(c, Coord{X: 0, Y: i})
	}

	return c
}

func szenario2(n int) []Coord {
	c := make([]Coord, n)

	for i := 0; i < n; i++ {
		c = append(c, Coord{X: n - 1, Y: i})
	}

	return c
}

func szenario3(n int) []Coord {
	c := make([]Coord, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				c = append(c, Coord{X: i, Y: j})
			}
		}
	}

	return c
}

func stopwatch(n int, c []Coord, b [][]int) {
	start := time.Now()
	// checkWinCoords(n, c)
	checkWinBoard(b)
	end := time.Now()
	// fmt.Println("checkWinCoords ,", n, ",", float64(end.Sub(start).Nanoseconds())/1000.0)
	fmt.Println("checkWinCoords ,", n, ",", float64(end.Sub(start).Microseconds())/1000.0)
}
