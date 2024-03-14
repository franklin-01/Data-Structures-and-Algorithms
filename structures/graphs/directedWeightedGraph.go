package graphs

func NewDirectedWeightedGraph(size int) {

	gph := &Graph{
		adjMatrix:  make([][]int, size),
		vertexData: make([]*byte, size),
		size:       &size,
	}

	for i := range gph.adjMatrix {
		gph.adjMatrix[i] = make([]int, size)
	}
}

func AddEdgeDWG(gph *Graph, vOne, vTwo, weight int) {
	if (vOne >= 0 && vOne < *gph.size) && (vTwo >= 0 && vTwo < *gph.size) {
		gph.adjMatrix[vOne][vTwo] = weight
	}
}
