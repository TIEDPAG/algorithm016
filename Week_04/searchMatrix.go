package t74

func searchMatrix(matrix [][]int, target int) bool {
	rowNum := searchRow(matrix, target)
	if rowNum < 0 {
		return false
	}
	return searchValue(matrix[rowNum], target)
}

func searchRow(matrix [][]int, target int) int {
	rowNum := len(matrix)
	if rowNum == 0 {
		return -1
	}
	columnNum := len(matrix[0])
	if columnNum == 0 {
		return -1
	}

	if target > matrix[rowNum-1][columnNum-1] ||
		target < matrix[0][0] {
		return -1
	}

	lowRow := 0
	for lowRow < rowNum {
		mid := (rowNum + lowRow) / 2
		if matrix[mid][columnNum-1] <= target && matrix[mid][0] >= target {
			return mid
		}
		if matrix[mid][columnNum-1] < target {
			lowRow = mid + 1
			continue
		}
		rowNum = mid
	}
	return lowRow
}

func searchValue(arr []int, target int) bool {
	l := len(arr)
	low := 0
	for low < l {
		mid := (low + l) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			l = mid
			continue
		}
		low = mid + 1
	}
	return false
}
