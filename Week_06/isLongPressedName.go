func isLongPressedName(name string, typed string) bool {
    l1, l2 := len(name), len(typed)
    i, j := 0, 0
    for j < l2 {
        if i < l1 && name[i] == typed[j] {
            i++
            j++
        } else if j > 0 && typed[j] == typed[j-1] {
            j++
        } else {
            return false
        }
    }
    return i == l1
}