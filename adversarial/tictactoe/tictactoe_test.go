package tictactoe

import (
	"testing"
)

func TestNextStates(t *testing.T) {
	ticTacToe := NewState(CellX)
	ticTacToe.Set(0, 0, CellX)
	ticTacToe.Set(1, 0, CellO)
	ticTacToe.Set(2, 0, CellX)
	ticTacToe.Set(0, 1, CellO)
	ticTacToe.Set(1, 1, CellO)
	ticTacToe.Set(2, 1, CellX)

	t.Log(ticTacToe.String())

	nextStates := ticTacToe.NextStates()
	if len(nextStates) != 3 {
		t.Fatal(len(nextStates))
	}
	//t.Log(nextStates)
}

func TestMinimax(t *testing.T) {
	ticTacToe := NewState(CellX)
	ticTacToe.Set(0, 0, CellX)
	ticTacToe.Set(1, 0, CellO)
	ticTacToe.Set(2, 0, CellX)
	ticTacToe.Set(0, 1, CellO)
	ticTacToe.Set(1, 1, CellO)
	ticTacToe.Set(2, 1, CellX)

	//t.Log(ticTacToe.String())
	// X │ O │ X
	//───┼───┼───
	// O │ O │ X
	//───┼───┼───
	//   │   │
	// best move for X should be 8 (2,2)
	// best move for O should be 7 (1,2)

	t.Run("Best move for X", func(t *testing.T) {
		bestMove := ticTacToe.BestMove()
		if bestMove != 8 {
			t.Fatalf("Expected %v got %v", width*2+2, bestMove)
		}
	})

	t.Run("Best move for O", func(t *testing.T) {
		ticTacToe.Current = CellO
		ticTacToe.Player = CellO
		bestMove := ticTacToe.BestMove()
		if bestMove != 7 {
			ticTacToe.board[bestMove] = CellO
			t.Log(ticTacToe.String())
			t.Fatalf("Expected %v got %v", 7, bestMove)
		}
	})
}
