# graph

Graph (and Tree) algorithms in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/jakecoffman/graph.svg)](https://pkg.go.dev/github.com/jakecoffman/graph)
![Build](https://github.com/jakecoffman/graph/actions/workflows/go.yml/badge.svg?branch=master)

Featuring:
- A* (Astar)
- Breadth First Search
- Depth First Search
- Uniform Cost Search (UCS or Dijkstra’s Algorithm)
- Beam search
- Chokudai search
- Monte Carlo Tree Search (MCTS)

Coming soon:
- Minimax

## Usage

This is not a library, just copy and paste what you need into your own code.

## Pathfinding

The `pathfinding` directory contains algorithms that find the shortest path from
start to goal. 

## Optimization

The `optimization` directory contains algorithms that find an optimal solution
to a problem that doesn't have a single clear goal.

## Bitset

The `bitset` directory contains functions wrapping common bitwise operations.

## Go performance tips

- Don't use `for _, copy := range`, use `for i := range` or `for i := 0; i < len(things); i++`.
- Don't use `map`, use arrays or slices
- Pool objects by making a big array `var pool = make([]Thing, 1_000_000)`, grab items like `thing := &pool[cursor]; cursor++`
- Turn off Garbage Collection once most things are pooled `debug.SetGCPercent(-1)`
- Once GC is off, prefer objects on the stack (`Thing{}`) not the heap (`&Thing{}`)
- Use the built-in benchmark and profiling functionality to find slow spots
- For even more performance turn arrays into bitsets. If multiple values are possible on the same position, then use multiple uints as "layers" of the grid.
