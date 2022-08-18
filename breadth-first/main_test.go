package main

import "testing"

type breadthTest struct {
	graph           map[string][]graphData
	name            string
	productToSearch string
	expected        string
}

var graph = make(map[string][]graphData)

func TestBreadthSearch(t *testing.T) {
	graph["you"] = []graphData{
		{"alice", "iphones"},
		{"bob", "burgers"},
		{"claire", "laptops"},
	}
	graph["alice"] = []graphData{
		{"peggy", "shorts"},
	}
	graph["bob"] = []graphData{
		{"anuj", "androids"},
		{"peggy", "shorts"},
	}
	graph["claire"] = []graphData{
		{"thomas", "water"},
		{"john", "records"},
	}
	graph["anuj"] = []graphData{}
	graph["peggy"] = []graphData{}
	graph["thomas"] = []graphData{}
	graph["john"] = []graphData{}

	var tests = []breadthTest{
		{graph, "you", "water", "thomas"},
		{graph, "you", "iphones", "alice"},
		{graph, "alice", "shorts", "peggy"},
		{graph, "bob", "androids", "anuj"},
		{graph, "claire", "androids", ""},
	}

	for _, test := range tests {
		output := BreadthSearch(test.graph, test.name, test.productToSearch)
		if output != test.expected {
			t.Errorf("Want %v, Got %v", test.expected, output)
		}
	}
}
