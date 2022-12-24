package pathfinding

import (
	"github.com/jakecoffman/graph/maze"
	"strings"
	"testing"
)

func Test_AstarArea(t *testing.T) {
	tests := []struct {
		start    string
		expected string
	}{
		{
			start: `
########
#     G#
#S #  G#
#     G#
########`,
			expected: `
########
#     G#
#Sx#  G#
# xxxxx#
########`,
		},
		{
			start: `
#########
#       #
# G S  G#
#       #
#########`,
			expected: `
#########
#       #
# xxS  G#
#       #
#########`,
		},
		{
			start: `
#########
#       #
# GwS  G#
#       #
#########`,
			expected: `
#########
#       #
# GwSxxx#
#       #
#########`,
		},
		{
			start: `
#########
##     G#
#G# S   #
##      #
#########`,
			expected: `
#########
##   xxx#
#G# Sx  #
##      #
#########`,
		},
	}

	for _, test := range tests {
		world := maze.NewWorld(test.start)
		start := world.FindOne(maze.Start)
		goals := world.FindAll(maze.Goal)

		path, found := AstarArea(start, goals)
		if len(test.expected) > 0 && !found {
			t.Error("Expected path not found")
		}
		if !found {
			continue
		}
		if world.RenderPath(path) != strings.TrimSpace(test.expected) {
			t.Error(world.RenderPath(path))
		}
	}
}
