package pathfinding

import (
	"strings"
	"testing"
)

func Test_Astar(t *testing.T) {
	tests := []struct {
		start    string
		expected string
	}{
		{
			start: `
########
#    #G#
#  # # #
#S#    #
########`,
			expected: `
########
#xxxx#x#
#x #x#x#
#S# xxx#
########`,
		}, {
			start: `
########
#      #
#  SG  #
#      #
########`,
			expected: `
########
#      #
#  Sx  #
#      #
########`,
		}, {
			start: `
########
#      #
#S ww G#
#  w   #
########`,
			expected: `
########
# xxxx #
#Sxwwxx#
#  w   #
########`,
		}, {
			start: `
########
#     ##
#S   #G#
#     ##
########
`,
			expected: ``, // no path
		},
	}

	for _, test := range tests {
		world := NewWorld(test.start)
		start := world.FindOne(Start)
		goal := world.FindOne(Goal)

		path, found := Astar(start, goal)
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
