func maxSumSubmatrix(matrix [][]int, k int) int {
    rows, cols := len(matrix), len(matrix[0])

    max := math.MinInt64
    for l := 0; l < cols; l++ {
        rowSums := make([]int, rows)
        for r := l; r < cols; r++ {
            for i := 0; i < rows; i++ {
                rowSums[i] += matrix[i][r]
            }

            _max := findMaxK(rowSums, k)
            // fmt.Printf("%v, max: %d\n", rowSums, _max)
            if _max == k {
                return _max
            }
            if _max > max {
                max = _max
            }
        }
    }
    return max
}

func findMaxK(rowSums []int, k int) int {
    max, dp, l := rowSums[0], rowSums[0], len(rowSums)
    for i := 1; i < l; i++ {
        if dp > 0 {
            dp += rowSums[i]
        } else {
            dp = rowSums[i]
        }
        if dp > max {
            max = dp
        }
    }

    if max <= k {
        return max
    }

    max = math.MinInt64
    for i := 0; i < l; i++ {
        sum := 0
        for j := i; j < l; j++ {
            sum += rowSums[j]
            if sum == k {
                return sum
            }

            if sum > max && sum < k {
                max = sum
            }
        }
    }
    return max
}