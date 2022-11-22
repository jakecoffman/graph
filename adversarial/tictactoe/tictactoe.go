package tictactoe

import (
	"fmt"
	"github.com/jakecoffman/graph/adversarial"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Cell int

func (c Cell) String() string {
	if c == CellBlank {
		return " "
	}
	if c == CellX {
		return "X"
	}
	return "O"
}

const (
	CellBlank = 0
	CellX     = 1
	CellO     = -1
)

const (
	width  = 3
	height = 3
)

type State struct {
	board []Cell
	turn  Cell
}

func NewState() *State {
	return &State{
		board: make([]Cell, width*height),
	}
}

func (s *State) String() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("\n %v │ %v │ %v \n", s.board[0], s.board[1], s.board[2]))
	str.WriteString("───┼───┼───\n")
	str.WriteString(fmt.Sprintf(" %v │ %v │ %v \n", s.board[3], s.board[4], s.board[5]))
	str.WriteString("───┼───┼───\n")
	str.WriteString(fmt.Sprintf(" %v │ %v │ %v ", s.board[6], s.board[7], s.board[8]))
	return str.String()
}

func (s *State) At(x, y int) Cell {
	return s.board[width*y+x]
}

func (s *State) Set(x, y int, cell Cell) {
	s.board[width*y+x] = cell
}

func (s *State) Index(i int) Cell {
	return s.board[i]
}

func (s *State) Play(index, color int) {
	if s.board[index] != CellBlank {
		log.Panicln("Illegal move", index)
	}
	s.board[index] = Cell(color)
	s.turn = -s.turn
}

func (s *State) Undo(index, color int) {
	s.board[index] = CellBlank
}

func (s *State) IsGameOver() bool {
	score := s.Score(1)
	if score != 0 {
		return true
	}
	var freeCellsLeft int
	for i := range s.board {
		if s.board[i] == CellBlank {
			freeCellsLeft++
		}
	}
	return freeCellsLeft == 0
}

func (s *State) Score(color int) int {
	player := Cell(color)

	if (s.board[0] == player && s.board[1] == player && s.board[2] == player) ||
		(s.board[3] == player && s.board[4] == player && s.board[5] == player) ||
		(s.board[6] == player && s.board[7] == player && s.board[8] == player) ||
		(s.board[0] == player && s.board[3] == player && s.board[6] == player) ||
		(s.board[1] == player && s.board[4] == player && s.board[7] == player) ||
		(s.board[2] == player && s.board[5] == player && s.board[8] == player) ||
		(s.board[0] == player && s.board[4] == player && s.board[8] == player) ||
		(s.board[2] == player && s.board[4] == player && s.board[6] == player) {
		return 1
	}

	opp := -player

	if (s.board[0] == opp && s.board[1] == opp && s.board[2] == opp) ||
		(s.board[3] == opp && s.board[4] == opp && s.board[5] == opp) ||
		(s.board[6] == opp && s.board[7] == opp && s.board[8] == opp) ||
		(s.board[0] == opp && s.board[3] == opp && s.board[6] == opp) ||
		(s.board[1] == opp && s.board[4] == opp && s.board[7] == opp) ||
		(s.board[2] == opp && s.board[5] == opp && s.board[8] == opp) ||
		(s.board[0] == opp && s.board[4] == opp && s.board[8] == opp) ||
		(s.board[2] == opp && s.board[4] == opp && s.board[6] == opp) {
		return -1
	}

	return 0
}

func (s *State) NextMoves() []int {
	var newStates []int
	for i, cell := range s.board {
		if cell == CellBlank {
			newStates = append(newStates, i)
		}
	}
	return newStates
}

func (s *State) BestMove(color int) int {
	// iterative deepening (not needed for TicTacToe, here for learning purposes)
	start := time.Now()
	bestMove := -1
	bestValue := math.MinInt64
	for distance := 1; distance < 100 && time.Now().Sub(start) < 100*time.Millisecond; distance++ {
		move, value := adversarial.Negamax(s, 100, color)
		if value > bestValue {
			bestValue = value
			bestMove = move
		}
	}

	return bestMove
}

var ZobristTable [3][3][2]uint64
var currentIsX uint64

func init() {
	rand.Seed(time.Now().UnixNano())

	currentIsX = rand.Uint64()

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			// p represents the player X or O
			for p := 0; p < 2; p++ {
				ZobristTable[x][y][p] = rand.Uint64()
			}
		}
	}
}

func (s *State) Hash(color int) uint64 {
	// TODO update this incrementally instead of generating from scratch each time
	var hash uint64

	if color == CellX {
		hash ^= currentIsX
	}

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			cell := s.At(x, y)
			if cell == CellX {
				hash ^= ZobristTable[x][y][0]
			} else if cell == CellO {
				hash ^= ZobristTable[x][y][1]
			}
		}
	}

	return hash
}
