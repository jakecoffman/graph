package adversarial

import "math"

type GameState interface {
	IsGameOver() bool
	Score(maximizingPlayer bool) int
	NextStates(maximizingPlayer bool) []GameState
}

// Minimax implements depth-limited minimax with alpha-beta pruning.
// maximizingPlayer should be true if the player goes first (X in TicTacToe or White in chess)
func Minimax(root GameState, maxDepth int, maximizingPlayer bool) int {
	return minimax(root, maxDepth, math.MinInt64, math.MaxInt64, maximizingPlayer)
}

// depth is the maximum depth to examine
// alpha represents the minimum score that the maximizing player is assured of
// beta represents the maximum score that the minimizing player is assured of
// Initially, alpha should be set to negative infinity and beta positive infinity.
func minimax(node GameState, depth, alpha, beta int, maximizingPlayer bool) int {
	if depth == 0 || node.IsGameOver() {
		return node.Score(maximizingPlayer)
	}
	if maximizingPlayer {
		value := math.MinInt64
		children := node.NextStates(maximizingPlayer)
		for i := range children {
			child := children[i]
			value = max(value, minimax(child, depth-1, alpha, beta, false))
			if value >= beta {
				break // beta pruning
			}
			alpha = max(alpha, value)
		}
		return value
	} else {
		value := math.MaxInt64
		children := node.NextStates(maximizingPlayer)
		for i := range children {
			child := children[i]
			value = min(value, minimax(child, depth-1, alpha, beta, true))
			if value <= alpha {
				break // alpha pruning
			}
			beta = min(beta, value)
		}
		return value
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
