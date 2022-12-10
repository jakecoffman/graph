package maze

import (
	"fmt"
	. "github.com/jakecoffman/graph"
	"strings"
)

type World struct {
	Width, Height int
	Map           []Node
}

// NewWorld is the World constructor. Takes a serialized World as input.
func NewWorld(input string) *World {
	str := strings.TrimSpace(input)
	rows := strings.Split(str, "\n")

	g := &World{
		Width:  len(rows[0]),
		Height: len(rows),
	}
	g.Map = make([]Node, g.Width*g.Height)

	for y, row := range rows {
		for x, raw := range row {
			kind, ok := Kinds[raw]
			if !ok {
				panic("Unknown rune: " + string(raw))
			} else {
				node := g.At(x, y)
				node.Kind = kind
				node.Pos = Pos{x, y}
			}
		}
	}

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			node := g.At(x, y)
			if node.Kind == Wall {
				continue
			}
			WalkNeighbors(Pos{x, y}, func(nX, nY int) {
				if nX >= g.Width {
					nX -= g.Width
				}
				if nX < 0 {
					nX += g.Width
				}
				if nY < 0 || nY >= g.Height {
					return
				}
				neighbor := g.At(nX, nY)
				if neighbor.Kind != Wall {
					node.Neighbors = append(node.Neighbors, neighbor)
				}
			})
		}
	}

	return g
}

// RenderPath serializes a path in a human readable way.
func (w *World) RenderPath(path []*Node) string {
	pathLocs := map[Pos]bool{}
	for _, p := range path {
		pathLocs[p.Pos] = true
	}
	rows := make([]string, w.Height)
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			t := w.At(x, y)
			r := ' '
			if pathLocs[Pos{x, y}] {
				r = 'x'
			} else if t != nil {
				r = Symbols[t.Kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

func (w *World) At(x, y int) *Node {
	return &w.Map[w.Width*y+x]
}

func (w *World) FindOne(kind Kind) *Node {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			node := w.At(x, y)
			if node.Kind == kind {
				return node
			}
		}
	}
	return nil
}

func (w *World) FindAll(kind Kind) (nodes []*Node) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			node := w.At(x, y)
			if node.Kind == kind {
				nodes = append(nodes, node)
			}
		}
	}
	return
}

type Node struct {
	Kind
	Pos
	Neighbors []*Node
	Visited   int
}

func (n *Node) String() string {
	return fmt.Sprintf(`[Node %v,%v %v]`, n.X, n.Y, Symbols[n.Kind])
}

type Kind rune

const (
	Plain = Kind(iota)
	Wall
	Water
	Start
	Goal
)

var Kinds = map[rune]Kind{
	' ': Plain,
	'#': Wall,
	'w': Water,
	'S': Start,
	'G': Goal,
}

var Symbols = map[Kind]rune{
	Plain: ' ',
	Wall:  '#',
	Water: 'w',
	Start: 'S',
	Goal:  'G',
}

var Costs = map[Kind]int{
	Plain: 1,
	Water: 5,
}
