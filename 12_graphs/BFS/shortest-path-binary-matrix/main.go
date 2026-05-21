package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {

	n := len(grid)

	// Invalid start or end
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	// 8 directions
	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	// Queue: row, col, distance
	queue := [][]int{{0, 0, 1}}

	// Mark visited
	grid[0][0] = 1

	// BFS
	for len(queue) > 0 {

		// Pop front
		current := queue[0]
		queue = queue[1:]

		row := current[0]
		col := current[1]
		distance := current[2]

		// Reached destination
		if row == n-1 && col == n-1 {
			return distance
		}

		// Explore neighbors
		for _, dir := range directions {

			newRow := row + dir[0]
			newCol := col + dir[1]

			// Boundary + valid cell check
			if newRow >= 0 &&
				newRow < n &&
				newCol >= 0 &&
				newCol < n &&
				grid[newRow][newCol] == 0 {

				// Mark visited
				grid[newRow][newCol] = 1

				// Add to queue
				queue = append(queue,
					[]int{newRow, newCol, distance + 1})
			}
		}
	}

	return -1
}

func main() {

	grid := [][]int{
		{0, 1},
		{1, 0},
	}

	fmt.Println(shortestPathBinaryMatrix(grid))
}
