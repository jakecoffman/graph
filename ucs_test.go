package pathfinding

import (
	"fmt"
	"testing"
)

func Test_UCS1(t *testing.T) {
	world := NewWorld(`
########
#    #G#
#  # # #
#S#    #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path := UCS(start, goal)
	if len(path) != 11 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_UCS2(t *testing.T) {
	world := NewWorld(`
########
#      #
#  SG  #
#      #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path := UCS(start, goal)
	if len(path) != 1 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_UCS3(t *testing.T) {
	world := NewWorld(`
########
#      #
#S ww G#
#  w   #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path := UCS(start, goal)
	if len(path) != 7 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}
