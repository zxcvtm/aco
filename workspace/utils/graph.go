package utils

import (
	"bufio"
	"os"
	"strings"
	"aco/workspace/schemas"
	"strconv"
)

func ReadGraph () [][]schemas.Edge {
	var graph [][]schemas.Edge
	file, err := os.Open("test-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var nodeEdges[]schemas.Edge

		data := scanner.Text()
		result := strings.Split(data, ",")
		for _, dist := range result {
			var edge schemas.Edge
			edge.Distance,_ = strconv.ParseFloat(dist, 64)
			nodeEdges = append(nodeEdges, edge)
		}
		graph = append(graph, nodeEdges)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return graph
}
