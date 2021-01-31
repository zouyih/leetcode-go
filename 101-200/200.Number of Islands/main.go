package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	m := len(grid)
	if m == 0 {
		return res
	}
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				res++
				dfs(grid, visited, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, visited [][]bool, i, j int) [][]bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || visited[i][j] {
		return visited
	}

	if grid[i][j] == '0' {
		return visited
	}

	visited[i][j] = true

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, direction := range directions {
		nextI, nextJ := i+direction[0], j+direction[1]
		visited = dfs(grid, visited, nextI, nextJ)
	}
	return visited
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
