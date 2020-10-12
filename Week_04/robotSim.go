package t874

func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	cur := 0
	xParam, yParam := []int{0, -1, 0, 1}, []int{1, 0, -1, 0}
	oMap:= make(map[int64]bool)
	for _, xy := range obstacles {
		ox := xy[0] + 30000
		oy := xy[1] + 30000
		k := (int64(ox) << 16) + int64(oy)
		oMap[k] = true
	}

	rs := 0
	for _, cmd := range commands {
		switch cmd {
		case -2:
			cur = (cur + 1) % 4
		case -1:
			cur = (cur + 3) % 4
		default:
			for i := 0; i < cmd; i++ {
				newX, newY := x + xParam[cur], y + yParam[cur]
				k := ((int64(newX) + 30000) << 16) + (int64(newY) + 30000)
				if oMap[k] {
					break
				}
				x, y = newX, newY

				rs = max(rs, x*x + y*y)
			}
		}
	}

	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
