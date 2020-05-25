package pathfinding

func WalkNeighbors(p Pos, callback func(x, y int)) {
	directions := []Pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, d := range directions {
		x, y := p.X+d.X, p.Y+d.Y
		callback(x, y)
	}
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ManhattanDistance(a, b Pos) int {
	dv := Abs(a.Y - b.Y)
	//dh := Min(Abs(a.X-b.X), Min(a.X+width-b.x, b.x+width-a.x)) wrap around
	dh := Abs(a.X - b.X)
	return dh + dv
}

type Queue struct {
	queue []*Node
}

func (q *Queue) Put(n *Node) *Queue {
	q.queue = append(q.queue, n)
	return q
}

func (q *Queue) Get() *Node {
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}
