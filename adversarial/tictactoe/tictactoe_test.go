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

func TestMinimax_Endgame(t *testing.T) {
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

func TestMinimax_Block(t *testing.T) {
	ticTacToe := NewState(CellX)
	ticTacToe = ticTacToe.Play(0)
	//t.Log(ticTacToe.String())
	// X │   │
	//───┼───┼───
	//   │   │
	//───┼───┼───
	//   │   │
	// best move for O should be 4 (1,1)

	// This test fails, it should go center but does not.
	t.Run("O should block X", func(t *testing.T) {
		ticTacToe.Player = CellO
		bestMove := ticTacToe.BestMove()
		if bestMove != 4 {
			ticTacToe.board[bestMove] = CellO
			t.Log(ticTacToe.String())
			t.Fatalf("Expected %v got %v", 4, bestMove)
		}
	})
}

func TestMinimax_Every_Move(t *testing.T) {
	for i := 0; i < 9; i++ {
		game := NewState(CellX)
		game = game.Play(i)

		for !game.IsGameOver() {
			if game.Player == CellX {
				game.Player = CellO
			} else {
				game.Player = CellX
			}
			bestMove := game.BestMove()
			game = game.Play(bestMove)
		}

		// every game should draw
		if game.Score() != 0 {
			t.Error("Failed to tie", i)
		}
	}
}
