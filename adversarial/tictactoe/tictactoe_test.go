package tictactoe

import (
	"math/rand"
	"testing"
	"time"
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

	nextMoves := ticTacToe.NextMoves()
	if len(nextMoves) != 3 {
		t.Fatal(len(nextMoves))
	}
	//t.Log(nextMoves)
}

func TestMinimax_Endgame(t *testing.T) {
	ticTacToe := NewState()
	ticTacToe.Play(0, CellX)
	ticTacToe.Play(1, CellO)
	ticTacToe.Play(2, CellX)
	ticTacToe.Play(3, CellO)
	ticTacToe.Play(5, CellX)
	ticTacToe.Play(4, CellO)
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
	ticTacToe.Play(0, CellX)
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMinimax_Every_Move(t *testing.T) {
	botPlayer := CellX

	for i := 0; i < 100_000; i++ {
		game := NewState()
		turn := CellX

		for !game.IsGameOver() {
			if turn == botPlayer {
				bestMove := game.BestMove(botPlayer)
				game.Play(bestMove, turn)
			} else {
				// non-bot player plays randomly
				nextMoves := game.NextMoves()
				game.Play(nextMoves[rand.Intn(len(nextMoves))], turn)
			}
			turn = -turn
		}

		// bot should always win or draw
		if game.Score(botPlayer) < 0 {
			t.Errorf("Bot player %v lost", colorToString(botPlayer))
			t.Log(game.String())
			t.Fatal()
		}
		// switch the bot player
		botPlayer = -botPlayer
	}
}

func colorToString(color int) string {
	if color == 1 {
		return "X"
	}
	return "O"
}
