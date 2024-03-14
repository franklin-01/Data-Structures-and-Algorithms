package graphs

import "fmt"

type Graph struct {
	adjMatrix  [][]int
	vertexData []*byte
	size       *int
	parent     []int
}

func AddVertexData(gph *Graph, vertex int, data byte) {
	if vertex >= 0 && vertex < *gph.size {
		gph.vertexData[vertex] = &data
	}
}

func PrintGraph(gph *Graph) {
	fmt.Println("Adjacecy Matrix:")
	fmt.Print("  ")
	for _, vertex := range gph.vertexData {
		fmt.Printf(" %c", *vertex)
	}

	fmt.Println("")

	for key, row := range gph.adjMatrix {
		fmt.Printf("%c ", *gph.vertexData[key])
		fmt.Println(row)
	}

	fmt.Println("")

	fmt.Println("Vertex data:")
	for key, vertex := range gph.vertexData {
		fmt.Printf("Vertex: %d: %c \n", key, *vertex)
	}

	fmt.Println("")

	fmt.Println("Connections for each vertex:")
	for vertexKey, vertex := range gph.vertexData {
		fmt.Printf("%c: ", *vertex)
		for key, value := range gph.adjMatrix[vertexKey] {
			if value != 0 {
				fmt.Printf("%c ", *gph.vertexData[key])
			}
		}
		fmt.Println()
	}

}

func dfsUtil(gph *Graph, vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%c ", *gph.vertexData[vertex])

	for i := range gph.adjMatrix {
		if (gph.adjMatrix[vertex][i] == 1) && (!visited[i]) {
			dfsUtil(gph, i, visited)
		}
	}
}

func Dfs(gph *Graph, startVertexData byte) {
	visited := make([]bool, *gph.size)
	for i := range visited {
		visited[i] = false
	}

	startVertex := VertexDataIndex(gph, startVertexData)
	dfsUtil(gph, startVertex, visited)
	fmt.Println()
}

func VertexDataIndex(gph *Graph, vertex byte) int {
	for i, gphVertex := range gph.vertexData {
		if vertex == *gphVertex {
			return i
		}
	}
	panic("This vertix does not exists.")
}

func Bfs(gph *Graph, startVertexData byte) {
	visited := make([]bool, *gph.size)
	queue := make([]int, *gph.size)
	queueStart, queueEnd := 0, 0
	startVertex := VertexDataIndex(gph, startVertexData)

	queue[queueEnd] = startVertex
	queueEnd++
	visited[startVertex] = true

	for queueStart < queueEnd {
		currentVertex := queue[queueStart]
		queueStart++
		fmt.Printf("%c ", *gph.vertexData[currentVertex])

		for i := range gph.adjMatrix[currentVertex] {
			if gph.adjMatrix[currentVertex][i] == 1 && !visited[i] {
				queue[queueEnd] = i
				queueEnd++
				visited[i] = true
			}
		}
	}

	fmt.Println()
}

func find(gph *Graph, i int) int {
	if gph.parent[i] == i {
		return i
	}
	return find(gph, gph.parent[i])
}

func onion(gph *Graph, x int, y int) {
	xRoot := find(gph, x)
	yRoot := find(gph, y)
	fmt.Printf("\n Union: %c + %c \n", *gph.vertexData[x], *gph.vertexData[y])
	gph.parent[xRoot] = yRoot
}

func UFCD(gph *Graph) bool {
	for i := range gph.adjMatrix {
		j := i + 1
		for range gph.adjMatrix {
			if gph.adjMatrix[i][j] == 1 {
				x := find(gph, i)
				y := find(gph, j)
				if x == y {
					return true
				}
				onion(gph, x, y)
			}
		}
	}
	return false
}
