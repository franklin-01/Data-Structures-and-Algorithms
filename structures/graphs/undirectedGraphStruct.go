package graphs

import "fmt"

func NewUndirectedGraph(size int) *Graph {

	gph := &Graph{
		adjMatrix:  make([][]int, size),
		vertexData: make([]*byte, size),
		size:       &size,
	}

	for i := range gph.adjMatrix {
		gph.adjMatrix[i] = make([]int, size)
	}

	return gph

}

func AddEdgeUG(gph *Graph, vOne, vTwo int) {
	if (vOne >= 0 && vOne < *gph.size) && (vTwo >= 0 && vTwo < *gph.size) {
		gph.adjMatrix[vOne][vTwo], gph.adjMatrix[vTwo][vOne] = 1, 1
	}
}

func dfsUtilUG(gph *Graph, vertex int, visited []bool, parent int) bool {
	visited[vertex] = true
	fmt.Printf("%c ", *gph.vertexData[vertex])

	for i := range gph.adjMatrix {
		if gph.adjMatrix[vertex][i] == 1 {
			if !visited[i] {
				result := dfsUtilUG(gph, i, visited, vertex)
				if result {
					return result
				}
			}
			if i != parent {
				return true
			}
		}
	}

	return false
}

func IsCyclicUG(gph *Graph) bool {
	visited := make([]bool, *gph.size)
	for i := range visited {
		visited[i] = false
	}
	for i := range visited {
		if !visited[i] {
			result := dfsUtilUG(gph, i, visited, -1)
			if result {
				return true
			}
		}
	}

	return false
}

func PrintUndirectedGraph(gph *Graph) {
	fmt.Println("Adjacecy Matrix:")
	fmt.Print("  ")
	for _, vertex := range gph.vertexData {
		fmt.Printf(" %c", *vertex)
	}

	fmt.Print("\n")
	for key, row := range gph.adjMatrix {
		fmt.Printf("%c ", *gph.vertexData[key])
		fmt.Println(row)
	}

	fmt.Println("vertex data:")
	for key, vertex := range gph.vertexData {
		fmt.Printf("Vertex: %d: %c \n", key, *vertex)
	}

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
