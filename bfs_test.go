package pathfinding

import (
	"log"
	"testing"
)

func TestWorld_BFS1(t *testing.T) {
	world := NewWorld(`
########
#    #G#
#  # # #
#S#    #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path := BFS(start, goal)
	if len(path) != 11 {
		t.Fatal(len(path))
	}
	log.Println(world.RenderPath(path))
}

func TestWorld_BFS2(t *testing.T) {
	world := NewWorld(`
########
#      #
#  SG  #
#      #
########
`)
	start := world.FindOne(Start)
	goal := world.FindOne(Goal)

	path := BFS(start, goal)
	if len(path) != 1 {
		t.Fatal(len(path))
	}
	log.Println(world.RenderPath(path))
}
