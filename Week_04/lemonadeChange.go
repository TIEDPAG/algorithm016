package t860

func lemonadeChange(bills []int) bool {
	moneys := []int{0, 0}
	for _, m := range bills {
		switch m {
		case 5:
			moneys[0]++
		case 10:
			if moneys[0] < 1 {
				return false
			}
			moneys[0]--
			moneys[1]++
		case 20:
			if moneys[1] > 0 && moneys[0] > 0 {
				moneys[0]--
				moneys[1]--
				break
			}
			if moneys[0] > 2 {
				moneys[0] -= 3
				break
			}
			return false
		}
	}
	return true
}
