func longestMountain(A []int) int {
    bCount, l := 0, len(A)
    left := 0
    for left < l-2 {
        right := left + 1
        if A[left] < A[right] {
            for right < l-1 && A[right] < A[right+1] {
                right++
            }
            if right < l-1 && A[right] > A[right+1] {
                for right < l-1 && A[right] > A[right+1] {
                    right++
                }
                if right-left+1 > bCount {
                    bCount = right-left+1
                }
            } else {
                right++
            }
        }
        left = right
    }
    return bCount
}