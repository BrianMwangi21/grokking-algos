package main

import (
	"fmt"
	"math"
)

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for node, cost := range costs {
		if _, inProcessed := processed[node]; cost < lowestCost && !inProcessed {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func Dijkstra(graph map[string]map[string]int, startNode string, finishNode string) int {
	costs := make(map[string]int)
	costs[finishNode] = math.MaxInt32

	parents := make(map[string]string)
	parents[finishNode] = ""

	processed := make(map[string]bool)

	for node, cost := range graph[startNode] {
		costs[node] = cost
		parents[node] = startNode
	}

	lowestCostNode := findLowestCostNode(costs, processed)

	for lowestCostNode != "" {
		for node, cost := range graph[lowestCostNode] {
			if _, isSet := costs[node]; !isSet {
				costs[node] = cost
			}
			
			newCost := costs[lowestCostNode] + cost

			if newCost < costs[node] {
				costs[node] = newCost
				parents[node] = lowestCostNode
			}
		}

		processed[lowestCostNode] = true
		lowestCostNode = findLowestCostNode(costs, processed)
	}

	return costs[finishNode]
}

func main() {
	fmt.Println("Hello, Dijkstra's algorithm")
}
