func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    if r == 0 {
        return nums[0]
    }
    if nums[l] < nums[r]{
        return nums[l]
    }

    for l <= r {
        mid := l + (r-l) / 2
        if nums[mid+1] < nums[mid] {
            return nums[mid+1]
        }

        if nums[mid-1] > nums[mid] {
            return nums[mid]
        }

        if nums[0] < nums[mid] {
            l = mid + 1
            continue
        }
        r = mid - 1
    }
    panic("error")
}