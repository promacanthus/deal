package graph

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	height := len(grid)
	width := len(grid[0])
	islands := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(grid, r, c, height, width)
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, r, c, h, w int) {
	if r < 0 || c < 0 || r >= h || c >= w || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c, h, w)
	dfs(grid, r+1, c, h, w)
	dfs(grid, r, c-1, h, w)
	dfs(grid, r, c+1, h, w)
}

func numIslandsBFS(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	height := len(grid)
	width := len(grid[0])
	islands := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '1' {
				islands++
				grid[r][c] = '0'
				queue := make([]int, 0)
				queue = append(queue, r*width+c)
				for len(queue) != 0 {
					id := queue[0]
					queue = queue[1:]
					row := id / width
					col := id % width
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, (row-1)*width+col)
						grid[row-1][col] = '0'
					}
					if row+1 < height && grid[row+1][col] == '1' {
						queue = append(queue, (row+1)*width+col)
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, row*width+col-1)
						grid[row][col-1] = '0'
					}
					if col+1 < width && grid[row][col+1] == '1' {
						queue = append(queue, row*width+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}
	return islands
}
