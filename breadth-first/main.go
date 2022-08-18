package main

import "fmt"

type graphData struct {
	name    string
	product string
}

func personSearched(name string, searched []string) bool {
	for _, n := range searched {
		if n == name {
			return true
		}
	}
	return false
}

func BreadthSearch(graph map[string][]graphData, name string, productToSearch string) string {
	var (
		result          string
		search_queue    []graphData
		searched        []string
		person, product string
	)

	search_queue = append(search_queue, graph[name]...)

	for len(search_queue) != 0 {
		person = search_queue[0].name
		product = search_queue[0].product
		search_queue = search_queue[1:]

		if !personSearched(person, searched) {
			if product == productToSearch {
				result = person
				break
			}

			search_queue = append(search_queue, graph[person]...)
			searched = append(searched, person)
		}
	}

	return result
}

func main() {
	fmt.Println("Hello, Breadth Search and I run at O(V+E)")
}
