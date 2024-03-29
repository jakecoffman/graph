package adversarial

import "math"

type GameState interface {
	// IsGameOver returns true when the game is over, win lose or draw.
	IsGameOver() bool
	// Evaluate returns a positive value if the color is winning, negative if losing.
	Evaluate(color int) int
	// NextMoves returns a list of all possible next moves.
	NextMoves() []int
	// Play takes a move and applies it to the state.
	Play(move, color int)
	// Undo rolls the state back before the move happened.
	Undo(move, color int)
	// Hash returns a hashable value for the game state so that
	// repeated states can be found by a map lookup. Google "Zobrist Hash".
	Hash(color int) uint64
}

const Inf = math.MaxInt64 - 1

// Negamax implements depth-limited negamax with alpha-beta pruning and transposition tables.
// It returns the index of the optimal state from the array generated by NextMoves.
// This assumes the player who needs to play next will be the maximizing player.
//
// depth is the maximum depth to examine
// alpha represents the minimum score that the maximizing player is assured of
// beta represents the maximum score that the minimizing player is assured of
// Initially, alpha should be set to negative infinity and beta positive infinity.
func Negamax[T GameState](state T, depth, alpha, beta int, color int) int {
	alphaOrig := alpha
	hash := state.Hash(color)
	ttEntry, ok := transpositionTable[hash]
	if ok && ttEntry.depth >= depth {
		if ttEntry.flag == flagExact {
			return ttEntry.value
		}
		if ttEntry.flag == flagLowerBound {
			alpha = max(alpha, ttEntry.value)
		} else if ttEntry.flag == flagUpperBound {
			beta = min(beta, ttEntry.value)
		}
		if alpha >= beta {
			return ttEntry.value
		}
	}

	if depth == 0 || state.IsGameOver() {
		return state.Evaluate(color)
	}
	value := math.MinInt64
	moves := state.NextMoves()
	for i := range moves {
		state.Play(moves[i], color)
		value = max(value, -Negamax(state, depth-1, -beta, -alpha, -color))
		state.Undo(moves[i], color)
		alpha = max(alpha, value)
		if alpha >= beta {
			break // cut-off
		}
	}

	ttEntry = entry{
		value: value,
		depth: depth,
	}
	if value <= alphaOrig {
		ttEntry.flag = flagUpperBound
	} else if value >= beta {
		ttEntry.flag = flagLowerBound
	} else {
		ttEntry.flag = flagExact
	}
	transpositionTable[hash] = ttEntry

	return value
}

var transpositionTable = map[uint64]entry{}

type flag int

const (
	flagExact = flag(iota)
	flagLowerBound
	flagUpperBound
)

type entry struct {
	flag
	depth int
	value int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
