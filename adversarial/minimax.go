package adversarial

import "math"

type GameState interface {
	IsGameOver() bool
	Score() int
	NextStates() []GameState
}

// Minimax implements depth-limited minimax with alpha-beta pruning.
// maximizingPlayer should be true if the player goes first (X in TicTacToe or White in chess)
func Minimax(state GameState, maxDepth int) int {
	move, _ := minimax(state, 0, maxDepth, math.MinInt64, math.MaxInt64, true)
	return move
}

// depth is the maximum depth to examine
// alpha represents the minimum score that the maximizing player is assured of
// beta represents the maximum score that the minimizing player is assured of
// Initially, alpha should be set to negative infinity and beta positive infinity.
func minimax(state GameState, depth, maxDepth, alpha, beta int, maximizingPlayer bool) (int, int) {
	if depth == maxDepth || state.IsGameOver() {
		return -1, state.Score()-depth
	}
	if maximizingPlayer {
		value := math.MinInt64
		move := -1
		children := state.NextStates()
		for i := range children {
			child := children[i]
			_, newValue := minimax(child, depth+1, maxDepth, alpha, beta, false)
			value = max(value, newValue)
			if value == newValue {
				move = i
			}
			if value >= beta {
				break // beta pruning
			}
			alpha = max(alpha, value)
		}
		return move, value
	} else {
		value := math.MaxInt64
		move := -1
		children := state.NextStates()
		for i := range children {
			child := children[i]
			_, newValue := minimax(child, depth+1, maxDepth, alpha, beta, true)
			value = min(value, newValue)
			if value == newValue {
				move = i
			}
			if value <= alpha {
				break // alpha pruning
			}
			beta = min(beta, value)
		}
		return move, value
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
