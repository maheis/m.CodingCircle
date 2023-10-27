package main

import tictactoe "github.com/maheis/CodingCircle/231025_TicTacToe/TicTacToe"

func main() {
	n := 3
	c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
}
