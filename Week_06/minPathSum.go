func minPathSum(grid [][]int) int {
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    cols := len(grid[0])
    if cols == 0 {
        return 0
    }

    for c := 1; c < cols; c++ {
        grid[0][c] += grid[0][c-1]
    }

    for r := 1; r < rows; r++ {
        grid[r][0] += grid[r-1][0]
        for c := 1; c < cols; c++ {
            grid[r][c] += min(grid[r][c-1], grid[r-1][c])
        }
    }
    return grid[rows-1][cols-1]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}