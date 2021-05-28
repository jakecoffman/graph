package optimization

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestBeam(t *testing.T) {
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	start := time.Now()
	const (
		beamSize = 200
		limit    = 100 * time.Millisecond
	)
	path := Beam(state, beamSize, limit)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
