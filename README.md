# graph

Graph (and Tree) algorithms in Go.

Featuring:
- A* (Astar)
- Breadth First Search
- Uniform Cost Search (UCS or Dijkstra’s Algorithm)
- Chokudai coming soon!

## Pathfinding

The `pathfinding` directory contains algorithms that find the shortest path from
start to goal. 

## Optimization

The `optimization` directory contains algorithms that find an optimal solution
to a problem that doesn't have a single clear goal.

## Go performance tips

- Don't use `range`, use `for i := 0; i < len(things); i++`
- Don't use `map`, use arrays or slices
- Pool objects by making a big array `var pool = make([]Thing, 1_000_000)`, grab items like `thing := &pool[cursor]; cursor++`
- Turn off Garbage Collection once most things are pooled `debug.SetGCPercent(-1)`
- Once GC is off, prefer objects on the stack (`Thing{}`) not the heap (`&Thing{}`)
- Use the built-in benchmark and profiling functionality to find slow spots
- For even more performance turn arrays into one or more `int`s and use bitwise operators to access the fields. If multiple values are possible on the same position, then use multiple ints as "layers" of the grid.
