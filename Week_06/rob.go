func rob(nums []int) int {
    if len(nums) < 2 {
        return nums[0]
    }
    return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }
    if l == 1 {
        return nums[0]
    }
    a, b := 0, nums[0]
    for i := 1; i < l; i++ {
        a += nums[i]
        a, b = b, max(a, b)
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}