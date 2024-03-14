package main

import (
	"fmt"
	"main/structures/graphs"
)

func main() {
	uGraph := graphs.NewUndirectedGraph(7)

	graphs.AddVertexData(uGraph, 0, 'A')
	graphs.AddVertexData(uGraph, 1, 'B')
	graphs.AddVertexData(uGraph, 2, 'C')
	graphs.AddVertexData(uGraph, 3, 'D')
	graphs.AddVertexData(uGraph, 4, 'E')
	graphs.AddVertexData(uGraph, 5, 'F')
	graphs.AddVertexData(uGraph, 6, 'G')

	graphs.AddEdgeUG(uGraph, 3, 0)
	graphs.AddEdgeUG(uGraph, 0, 2)
	graphs.AddEdgeUG(uGraph, 0, 3)
	graphs.AddEdgeUG(uGraph, 0, 4)
	graphs.AddEdgeUG(uGraph, 4, 2)
	graphs.AddEdgeUG(uGraph, 2, 5)
	graphs.AddEdgeUG(uGraph, 2, 1)
	graphs.AddEdgeUG(uGraph, 2, 6)
	graphs.AddEdgeUG(uGraph, 1, 5)

	fmt.Printf("\nDepth First Search starting from vertex D: \n")
	graphs.Dfs(uGraph, 'D')

	fmt.Printf("\nBreadth First Search starting from vertex D: \n")
	graphs.Bfs(uGraph, 'D')

	isCiclic := graphs.IsCyclicUG(uGraph)

	fmt.Printf("\nGraph has cycle: %v\n", isCiclic)
}
