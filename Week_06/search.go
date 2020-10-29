func search(nums []int, target int) int {
    n := len(nums)-1
    l, r := 0, n

    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] == target {
            return mid
        }

        if nums[0] <= nums[mid] {
            if nums[0] <= target && target < nums[mid] {
                r = mid - 1
                continue
            }
            l = mid + 1
            continue
        }
        if nums[mid] < target && target <= nums[n] {
            l = mid + 1
            continue
        }
        r = mid - 1
    }
    return -1
}