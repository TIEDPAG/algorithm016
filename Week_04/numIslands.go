package t200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	l, r := len(grid), len(grid[0])
	rs := 0
	for i := 0; i < r; i++ {
		for j := 0; j < l; j++ {
			if grid[j][i] == '1' {
				rs++
				dfs(grid, i, j)
			}
		}
	}
	return rs
}

func dfs(grid [][]byte, i, j int) {
	l, r := len(grid), len(grid[0])

	if i < 0 || j < 0 || i >= r || j >= l || grid[j][i] == '0' {
		return
	}

	grid[j][i] = '0'

	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
}
