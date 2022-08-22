package main

import "testing"

type dijkstraTest struct {
	graph      map[string]map[string]int
	startNode  string
	finishNode string
	expected   int
}

var graph = make(map[string]map[string]int)

func TestDijkstra(t *testing.T) {
	graph["start"] = map[string]int{}
	graph["start"]["b"] = 5
	graph["start"]["c"] = 2

	graph["b"] = map[string]int{}
	graph["b"]["d"] = 1
	graph["b"]["e"] = 4

	graph["c"] = map[string]int{}
	graph["c"]["b"] = 8
	graph["c"]["d"] = 7

	graph["d"] = map[string]int{}
	graph["d"]["stop"] = 1

	graph["e"] = map[string]int{}
	graph["e"]["d"] = 6
	graph["e"]["stop"] = 3
	
	graph["stop"] = map[string]int{}

	var tests = []dijkstraTest{
		{graph, "start", "stop", 7},
		{graph, "start", "e", 9},
		{graph, "start", "d", 6},
		{graph, "b", "stop", 2},
		{graph, "e", "stop", 3},
	}

	for _, test := range tests {
	 	output := Dijkstra(test.graph, test.startNode, test.finishNode)
	 	if output != test.expected {
			t.Errorf("Want %v, Got %v", test.expected, output)
		}
	}
}
