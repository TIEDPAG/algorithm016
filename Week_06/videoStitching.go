func videoStitching(clips [][]int, T int) int {
    arr := make([]int, T)
    l := len(clips)
    for i := 0; i < l; i++ {
        start, end := clips[i][0], clips[i][1]
        if start < T && arr[start] < end {
            arr[start] = end
        }
    }

    last, prev, step := arr[0], 0 , 0
    for i := 0; i < T; i++ {
        if arr[i] > last {
            last = arr[i]
        }
        if i == last {
            return -1
        }
        if i == prev {
            step++
            prev = last
        }
    }
    return step
}