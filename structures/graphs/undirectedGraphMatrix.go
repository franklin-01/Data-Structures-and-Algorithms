package graphs

import "fmt"

// An undirected graph with adjacency maxtrix

var vertexData = [4]byte{'A', 'B', 'C', 'D'}

var adjacency_matrix = [4][4]int{
	{0, 1, 1, 1}, // Edges for A
	{1, 0, 1, 0}, // Edges for B
	{1, 1, 0, 0}, // Edges for C
	{1, 0, 0, 0}, // Edges for D
}

func printAdjacecyMatrix(arr [4][4]int) {
	fmt.Println("Adjacency Matrix")
	fmt.Printf("  ")
	for _, char := range vertexData {
		fmt.Printf(" %c", char)
	}
	fmt.Print("\n")
	for key, row := range arr {
		fmt.Printf("%c ", vertexData[key])
		fmt.Println(row)
	}
}

func printConnectionsInAdjacecyMatrix(arr [4][4]int, vertices [4]byte) {
	fmt.Println("Connections for each vertex")
	for vertexKey, vertex := range vertices {
		fmt.Printf("%c: ", vertex)
		for key, value := range arr[vertexKey] {
			if value != 0 {
				fmt.Printf("%c ", vertices[key])
			}
		}
		fmt.Println()
	}
}

func UndirectedGraphMatrixRepresentation() {
	printAdjacecyMatrix(adjacency_matrix)
	printConnectionsInAdjacecyMatrix(adjacency_matrix, vertexData)
}
