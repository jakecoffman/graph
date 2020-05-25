package pathfinding

import (
	"fmt"
	"strings"
)

type World struct {
	width, height int
	world         [][]*Node
}

// don't you dare close your eyes
func NewWorld(input string) *World {
	str := strings.TrimSpace(input)
	rows := strings.Split(str, "\n")

	g := &World{
		width:  len(rows[0]),
		height: len(rows),
		world:  nil,
	}

	for x := 0; x < g.width; x++ {
		g.world = append(g.world, []*Node{})
		for y := 0; y < g.height; y++ {
			g.world[x] = append(g.world[x], &Node{Pos: Pos{x, y}})
		}
	}

	for y, row := range rows {
		for x, raw := range row {
			kind, ok := Kinds[raw]
			if !ok {
				panic("Unknown rune: " + string(raw))
			} else {
				g.At(x, y).Kind = kind
			}
		}
	}

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			node := g.At(x, y)
			if node.Kind == Wall {
				continue
			}
			WalkNeighbors(Pos{x, y}, func(nX, nY int) {
				if nX >= g.width {
					nX -= g.width
				}
				if nX < 0 {
					nX += g.width
				}
				if nY < 0 || nY >= g.height {
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

func (w *World) RenderPath(path []*Node) string {
	pathLocs := map[Pos]bool{}
	for _, p := range path {
		pathLocs[p.Pos] = true
	}
	rows := make([]string, w.height)
	for x := 0; x < w.width; x++ {
		for y := 0; y < w.height; y++ {
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
	return w.world[x][y]
}

func (w *World) FindOne(kind Kind) *Node {
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			node := w.At(x, y)
			if node.Kind == kind {
				return node
			}
		}
	}
	return nil
}

type Pos struct {
	X, Y int
}

func (p Pos) Add(other Pos) Pos {
	return Pos{p.X + other.X, p.Y + other.Y}
}

func (p Pos) Sub(other Pos) Pos {
	return Pos{other.X - p.X, other.Y + p.Y}
}

type Node struct {
	Kind
	Pos
	Neighbors []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf(`[Node %v,%v %v]`, n.X, n.Y, Symbols[n.Kind])
}

type Kind rune

const (
	Plain = Kind(iota)
	Wall
	Start
	Goal
)

var Kinds = map[rune]Kind{
	' ': Plain,
	'#': Wall,
	'S': Start,
	'G': Goal,
}

var Symbols = map[Kind]rune{
	Plain: ' ',
	Wall:  '#',
	Start: 'S',
	Goal:  'G',
}
