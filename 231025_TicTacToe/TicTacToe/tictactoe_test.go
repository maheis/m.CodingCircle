package tictactoe_test

import (
	"testing"

	tictactoe "github.com/maheis/CodingCircle/231025_TicTacToe/TicTacToe"
	"github.com/stretchr/testify/assert"
)

func TestTicTacToe(t *testing.T) {

	t.Run("CheckGame - no coordinates given", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrNoCoords)
	})

	t.Run("CheckGame - n must be greater than 0", func(t *testing.T) {
		n := 0
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrBoardToSmall)
	})

	t.Run("CheckGame - not enough coordinates", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 1, Y: 1}, {X: 2, Y: 2}}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrNotEnoughCoords)
	})

	t.Run("CheckGame - to much coordinates", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 0, Y: 0}}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrToMuchCoords)
	})

	t.Run("CheckGame - coordinates out of range +", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 3}}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrCoordsOutOfRange)
	})

	t.Run("CheckGame - coordinates out of range -", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: -2}}
		_, err := tictactoe.CheckGame(n, c)

		assert.ErrorIs(t, err, tictactoe.ErrCoordsOutOfRange)
	})

	t.Run("CheckGame - player won diagonal lr", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("CheckGame - player won diagonal rl", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 2, Y: 0}, {X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("CheckGame - player won vertical", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("CheckGame - player won horizontal", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("CheckGame - all fields", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("CheckGame - no winner", func(t *testing.T) {
		n := 3
		c := []tictactoe.Coord{{X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 2}}
		res, err := tictactoe.CheckGame(n, c)

		assert.NoError(t, err)
		assert.False(t, res)
	})
}
