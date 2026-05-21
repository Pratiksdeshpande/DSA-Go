package main

import "fmt"

func main() {

	grid := [][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '0', '1'},
		{'0', '1', '0', '0', '0'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println(numIslands(grid))
}

var (
	// Directions: up, down, left, right
	directions = [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
)

func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	islands := traverseGrid(grid, rows, cols)

	return islands
}

func traverseGrid(grid [][]byte, rows, cols int) int {
	islands := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if grid[r][c] == '1' {
				islands++

				// BFS Queue
				queue := [][]int{{r, c}}

				// Mark visited
				grid[r][c] = '0'

				for len(queue) > 0 {
					cell := queue[0]
					queue = queue[1:]

					row := cell[0]
					col := cell[1]

					for _, direction := range directions {
						newRow := row + direction[0]
						newCol := col + direction[1]

						if newRow >= 0 &&
							newRow < rows &&
							newCol >= 0 &&
							newCol < cols &&
							grid[newRow][newCol] == '1' {
							grid[newRow][newCol] = '0'
							queue = append(queue, []int{newRow, newCol})
						}
					}
				}

			}
		}
	}

	return islands
}
