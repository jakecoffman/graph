package pathfinding

import (
	"fmt"
	"testing"
)

func Test_DFS1(t *testing.T) {
	world := NewWorld(`
########
#    #G#
#  # # #
#S#    #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := DFS(start, goal)
	if !found || len(path) != 11 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_DFS2(t *testing.T) {
	world := NewWorld(`
########
#      #
#  SG  #
#      #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := DFS(start, goal)
	if !found || len(path) != 1 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}

func Test_DFS_NoPath(t *testing.T) {
	world := NewWorld(`
########
#     ##
#  S #G#
#     ##
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path, found := DFS(start, goal)
	if found || len(path) != 0 {
		t.Fatal(len(path))
	}
	fmt.Println(world.RenderPath(path))
}
