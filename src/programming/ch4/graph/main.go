//graph键类型为string,值类型为map[string]int,表示一个字符串集合，因此，graph建立了一个从字符串到字符串集合的映射
package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func main() {
	from, to := "dqwd", "hghfd"
	addEdge(from, to)
	fmt.Println(graph)
	fmt.Println(hasEdge(from, "fewrw"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
