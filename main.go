package main
import (
	// Standard
	"fmt"
	"aco/workspace/utils"
	"aco/workspace/schemas"
)

func main() {
	fmt.Println("its wokrs")
	fmt.Println("Lets start reading the graph file")

	graph  := utils.ReadGraph()
	garden := schemas.Garden{}
	garden = garden.InitGarden(graph, 0.5, 0.5)
	garden = garden.Iterate()
	fmt.Println(garden.GetBestRoute())

}