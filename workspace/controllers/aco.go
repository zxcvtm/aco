package controllers

import (
	"fmt"
	"time"
	"aco/workspace/utils"
	"aco/workspace/schemas"
)

func AcoAlgorithm(graphArray [][] float64) (response schemas.ApiResponse) {
	start := time.Now()

	graph  := utils.ArrayToEdges(graphArray)
	garden := schemas.Garden{}
	garden = garden.InitGarden(graph, 0.8, 0.4, 0.5, 0.4)

	bestCase := 1000000.0
	bestAnt := schemas.Ant{}
	for i:=0; i<=500; i++ {
		garden = garden.IterateAsync()
		antRoute := garden.GetBestRoute()
		if bestCase > antRoute.GetRouteDistante() {
			bestCase = antRoute.GetRouteDistante()
			bestAnt = antRoute
		}
		if (i%100)==0  {
			fmt.Println(bestCase)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Finished at ", elapsed, " best case:", bestCase)
	reply := make(map[string]interface{})
	reply["order"] = bestAnt.GetVisited()
	reply["elapsedTime"] = elapsed
	reply["Distance"] = bestCase
	return response.Success(reply, "Sucessfuly Result")
}
