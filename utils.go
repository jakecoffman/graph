package pathfinding

import "container/heap"

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

// Priority queue copy-pasted from Go docs
type Item struct {
	*Node        // The value of the item; arbitrary.
	Priority int // The Priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The Index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the Priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *Node, priority int) {
	item.Node = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func (pq PriorityQueue) Empty() bool {
	return len(pq) == 0
}
