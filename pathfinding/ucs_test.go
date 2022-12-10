package pathfinding

import (
	"fmt"
	"github.com/jakecoffman/graph/maze"
	"testing"
)

func Test_UCS1(t *testing.T) {
	world := maze.NewWorld(`
########
#    #G#
#  # # #
#S#    #
########
`)
	start := world.FindOne(maze.Start)
	goal := world.FindOne(maze.Goal)

	path, found := UCS(start, goal)
	if !found || len(path) != 11 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_UCS2(t *testing.T) {
	world := maze.NewWorld(`
########
#      #
#  SG  #
#      #
########
`)
	start := world.FindOne(maze.Start)
	goal := world.FindOne(maze.Goal)

	path, found := UCS(start, goal)
	if !found || len(path) != 1 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_UCS3(t *testing.T) {
	world := maze.NewWorld(`
########
#      #
#S ww G#
#  w   #
########
`)
	start := world.FindOne(maze.Start)
	goal := world.FindOne(maze.Goal)

	path, found := UCS(start, goal)
	if !found || len(path) != 7 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_UCS_NoPath(t *testing.T) {
	world := maze.NewWorld(`
########
#     ##
#S   #G#
#     ##
########
`)
	start := world.FindOne(maze.Start)
	goal := world.FindOne(maze.Goal)

	path, found := UCS(start, goal)
	if found || len(path) != 0 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}
