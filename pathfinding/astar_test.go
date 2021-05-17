package pathfinding

import (
	"fmt"
	"testing"
)

func Test_Astar1(t *testing.T) {
	world := NewWorld(`
########
#    #G#
#  # # #
#S#    #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := Astar(start, goal)
	if !found || len(path) != 11 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_Astar2(t *testing.T) {
	world := NewWorld(`
########
#      #
#  SG  #
#      #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := Astar(start, goal)
	if !found || len(path) != 1 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_Astar3(t *testing.T) {
	world := NewWorld(`
########
#      #
#S ww G#
#  w   #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := Astar(start, goal)
	if !found || len(path) != 7 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_Astar_NoPath(t *testing.T) {
	world := NewWorld(`
########
#     ##
#S   #G#
#     ##
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := Astar(start, goal)
	if found || len(path) != 0 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}
