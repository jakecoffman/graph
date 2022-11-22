package main

import (
	"bufio"
	"fmt"
	"github.com/jakecoffman/graph/adversarial/tictactoe"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	reader *bufio.Reader
}

func NewInput() *Input {
	return &Input{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (i *Input) Read() string {
	str, err := i.reader.ReadString('\n')
	check(err)
	return strings.Split(str, "\n")[0]
}

func main() {
	input := NewInput()

	fmt.Println("Do you want to go first? (Y|n)")
	str := input.Read()

	player := tictactoe.CellO
	if str == "" || strings.HasPrefix(strings.ToUpper(str), "Y") {
		player = tictactoe.CellX
	}

	state := tictactoe.NewState()
	turn := tictactoe.CellX

	for !state.IsGameOver() {
		if player == turn {
			fmt.Println(state)
			fmt.Println("Enter index to move: ")
			str = input.Read()
			i, err := strconv.Atoi(str)
			check(err)
			if state.Index(i) != tictactoe.CellBlank {
				continue
			}
			state.Play(i, turn)
		} else {
			move := state.BestMove(turn)
			state.Play(move, turn)
		}
		turn = -turn
	}

	fmt.Println(state)
	score := state.Score(turn)
	if score == 0 {
		fmt.Println("TIE!")
	} else if score < 0 {
		fmt.Println("LOSE")
	} else {
		fmt.Println("WIN")
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
