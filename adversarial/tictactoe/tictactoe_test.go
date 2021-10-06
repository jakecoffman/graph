package tictactoe

import (
	"testing"
)

func TestNextStates(t *testing.T) {
	ticTacToe := NewState()
	ticTacToe.Set(0, 0, CellX)
	ticTacToe.Set(1, 0, CellO)
	ticTacToe.Set(2, 0, CellX)
	ticTacToe.Set(0, 1, CellO)
	ticTacToe.Set(1, 1, CellO)
	ticTacToe.Set(2, 1, CellX)

	//t.Log(ticTacToe.String())

	nextStates := ticTacToe.NextStates(1)
	if len(nextStates) != 3 {
		t.Fatal(len(nextStates))
	}
	//t.Log(nextStates)
}

func TestMinimax_Endgame(t *testing.T) {
	ticTacToe := NewState()
	ticTacToe = ticTacToe.Play(0, CellX).
		Play(1, CellO).
		Play(2, CellX).
		Play(3, CellO).
		Play(5, CellX).
		Play(4, CellO)
	t.Log(ticTacToe)
	// X │ O │ X
	//───┼───┼───
	// O │ O │ X
	//───┼───┼───
	//   │   │
	// best move for X should be 8 (2,2)
	// best move for O should be 7 (1,2)

	t.Run("Best move for X", func(t *testing.T) {
		bestMove := ticTacToe.BestMove(CellX)
		if bestMove != 8 {
			t.Fatalf("Expected %v got %v", width*2+2, bestMove)
		}
	})

	t.Run("Best move for O", func(t *testing.T) {
		bestMove := ticTacToe.BestMove(CellO)
		if bestMove != 7 {
			ticTacToe.board[bestMove] = CellO
			t.Log(ticTacToe.String())
			t.Fatalf("Expected %v got %v", 7, bestMove)
		}
	})
}

func TestMinimax_Block(t *testing.T) {
	ticTacToe := NewState()
	ticTacToe = ticTacToe.Play(0, CellX)
	//t.Log(ticTacToe.String())
	// X │   │
	//───┼───┼───
	//   │   │
	//───┼───┼───
	//   │   │
	// best move for O should be 4 (1,1)

	// This test fails, it should go center but does not.
	t.Run("O should block X", func(t *testing.T) {
		bestMove := ticTacToe.BestMove(CellO)
		if bestMove != 4 {
			ticTacToe.board[bestMove] = CellO
			t.Log(ticTacToe.String())
			t.Fatalf("Expected %v got %v", 4, bestMove)
		}
	})
}

func TestMinimax_Every_Move(t *testing.T) {
	color := CellX

	for i := 0; i < 9; i++ {
		game := NewState()
		game = game.Play(i, color)

		for !game.IsGameOver() {
			color = -color
			if color == 1 {
				bestMove := game.BestMove(CellX)
				game = game.Play(bestMove, color)
			} else {
				bestMove := game.BestMove(CellO)
				game = game.Play(bestMove, color)
			}
		}

		// every game should draw
		if game.Score() != 0 {
			t.Error("Failed to tie", i)
		}
	}
}
