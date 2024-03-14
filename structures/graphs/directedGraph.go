package graphs

import "fmt"

func AddEdgeDG(gph *Graph, vOne int, vTwo int) {
	if (vOne >= 0 && vOne < *gph.size) && (vTwo >= 0 && vTwo < *gph.size) {
		gph.adjMatrix[vOne][vTwo] = 1
	}
}

func dsfUtilDG(gph *Graph, vertex int, visited []bool, recStack []bool) bool {
	visited[vertex] = true
	recStack[vertex] = true
	fmt.Printf("\nCurrent vertex: %c\n", *gph.vertexData[vertex])

	for i := range visited {
		if gph.adjMatrix[vertex][i] == 1 {
			if !visited[i] {
				result := dsfUtilDG(gph, i, visited, recStack)
				if result {
					return true
				}
			}
			if recStack[i] {
				return true
			}
		}
	}
	recStack[vertex] = false
	return false
}

func IsCycicDG(gph *Graph) bool {
	visited := make([]bool, *gph.size)
	recStack := make([]bool, *gph.size)
	for i := range visited {
		visited[i] = false
		recStack[i] = false
	}

	for i := range visited {
		if !visited[i] {
			fmt.Println()
			result := dsfUtilDG(gph, i, visited, recStack)
			if result {
				return true
			}
		}
	}

	return false
}
