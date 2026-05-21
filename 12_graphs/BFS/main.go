package main

import "fmt"

func main() {

	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0},
		3: {1},
		4: {1},
	}

	BFS(graph, 0)
}

func BFS(graph map[int][]int, start int) {

	// Track visited nodes
	visited := make(map[int]bool)

	// Queue for BFS
	queue := []int{}

	// Mark start node visited
	visited[start] = true

	// Add start node to queue
	queue = append(queue, start)

	// Continue until queue becomes empty
	for len(queue) > 0 {

		// Remove front node from queue
		node := queue[0]
		queue = queue[1:]

		// Print current node
		fmt.Print(node, " ")

		// Traverse neighbors
		for _, neighbor := range graph[node] {

			// Process only unvisited nodes
			if !visited[neighbor] {

				// Mark visited
				visited[neighbor] = true

				// Add to queue
				queue = append(queue, neighbor)
			}
		}
	}
}
