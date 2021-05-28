package optimization

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestBeam(t *testing.T) {
	world := NewWorld(`
################
#    #G   # G  #
#  # # #  # #  #
#S#  # #     # #
# # G     #  #G#
# # ##### #  # #
#     G        #
################
`)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	start := time.Now()
	path := Beam(3, state)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
