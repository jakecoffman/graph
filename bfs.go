package pathfinding

func BFS(start, goal *Node) []*Node {
	frontier := Queue{}
	frontier.Put(start)
	cameFrom := map[*Node]*Node{
		start: nil,
	}

	for !frontier.Empty() {
		current := frontier.Get()

		// early exit
		if current == goal {
			break
		}

		for _, next := range current.Neighbors {
			if _, ok := cameFrom[next]; !ok {
				frontier.Put(next)
				cameFrom[next] = current
			}
		}
	}

	current := goal
	var path []*Node
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
