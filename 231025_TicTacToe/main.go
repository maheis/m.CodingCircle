package main

import tictactoe "github.com/maheis/CodingCircle/231025_TicTacToe/TicTacToe"

func main() {
	var c []tictactoe.Coord

	n := 3
	//Aufgabe Diagonal \
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//H
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//V
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	// Diagonal /
	c = []tictactoe.Coord{{X: 2, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//Alle Felder
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//No Winner
	c = []tictactoe.Coord{{X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//no coordinates
	c = []tictactoe.Coord{}
	tictactoe.CheckGame(n, c)
	//n to small
	n = 0
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	n = 3
	//not enough coordinates
	c = []tictactoe.Coord{{X: 1, Y: 1}, {X: 2, Y: 2}}
	tictactoe.CheckGame(n, c)
	//to much coordinates
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 0, Y: 0}}
	tictactoe.CheckGame(n, c)
	//out of range
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 3}}
	tictactoe.CheckGame(n, c)
	//out of range
	c = []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: -2}}
	tictactoe.CheckGame(n, c)
}
