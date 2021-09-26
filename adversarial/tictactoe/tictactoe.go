package tictactoe

import (
	"fmt"
	"github.com/jakecoffman/graph/adversarial"
	"log"
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
	CellBlank = Cell(iota)
	CellX
	CellO
)

const (
	width  = 3
	height = 3
)

type State struct {
	board     []Cell
	// Current is the player that needs to make a move next
	Current   Cell
	// Player is the maximizer: they want to maximize their score
	Player    Cell
}

func NewState(player Cell) *State {
	return &State{
		board:     make([]Cell, width*height),
		Player:    player,
		Current:   CellX,
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

func (s *State) Clone() *State {
	newState := NewState(s.Player)
	copy(newState.board, s.board)
	newState.Current = s.Current
	return newState
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

func (s *State) Play(i int) *State {
	if s.board[i] != CellBlank {
		log.Panicln("Illegal move", i)
	}
	n := s.Clone()
	n.board[i] = n.Current
	if n.Current == CellX {
		n.Current = CellO
	} else {
		n.Current = CellX
	}
	return n
}

func (s *State) IsGameOver() bool {
	score := s.Score()
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

func (s *State) Score() int {
	x := (s.board[0] == CellX && s.board[1] == CellX && s.board[2] == CellX) ||
		(s.board[3] == CellX && s.board[4] == CellX && s.board[5] == CellX) ||
		(s.board[6] == CellX && s.board[7] == CellX && s.board[8] == CellX) ||
		(s.board[0] == CellX && s.board[3] == CellX && s.board[6] == CellX) ||
		(s.board[1] == CellX && s.board[4] == CellX && s.board[7] == CellX) ||
		(s.board[2] == CellX && s.board[5] == CellX && s.board[8] == CellX) ||
		(s.board[0] == CellX && s.board[4] == CellX && s.board[8] == CellX) ||
		(s.board[2] == CellX && s.board[4] == CellX && s.board[6] == CellX)

	o := (s.board[0] == CellO && s.board[1] == CellO && s.board[2] == CellO) ||
		(s.board[3] == CellO && s.board[4] == CellO && s.board[5] == CellO) ||
		(s.board[6] == CellO && s.board[7] == CellO && s.board[8] == CellO) ||
		(s.board[0] == CellO && s.board[3] == CellO && s.board[6] == CellO) ||
		(s.board[1] == CellO && s.board[4] == CellO && s.board[7] == CellO) ||
		(s.board[2] == CellO && s.board[5] == CellO && s.board[8] == CellO) ||
		(s.board[0] == CellO && s.board[4] == CellO && s.board[8] == CellO) ||
		(s.board[2] == CellO && s.board[4] == CellO && s.board[6] == CellO)

	if x {
		if s.Player == CellX {
			return 20
		} else {
			return -20
		}
	}
	if o {
		if s.Player == CellX {
			return -20
		} else {
			return 20
		}
	}
	return 0
}

func (s *State) NextStates() []adversarial.GameState {
	var newStates []adversarial.GameState
	for i, cell := range s.board {
		if cell == CellBlank {
			n := s.Play(i)
			newStates = append(newStates, n)
		}
	}
	return newStates
}

func (s *State) BestMove() int {
	// iterative deepening (not needed for TicTacToe, here for learning purposes)
	start := time.Now()
	var best int
	for distance := 1; distance < 1000 && time.Now().Sub(start) < 100*time.Millisecond; distance++ {
		best = adversarial.Negamax(s, 1000)
	}

	// now we must figure out what move this actually was
	var legalMoves []int
	for i := range s.board {
		if s.Index(i) == CellBlank {
			legalMoves = append(legalMoves, i)
		}
	}
	return legalMoves[best]
}

var ZobristTable [3][3][2]uint64
var currentIsX uint64
var playerIsX uint64

func init() {
	rand.Seed(time.Now().UnixNano())

	currentIsX = rand.Uint64()
	playerIsX = rand.Uint64()

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			// p represents the player X or O
			for p := 0; p < 2; p++ {
				ZobristTable[x][y][p] = rand.Uint64()
			}
		}
	}
}

func (s *State) Hash() uint64 {
	var hash uint64

	if s.Player == CellX {
		hash ^= playerIsX
	}
	if s.Current == CellX {
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
