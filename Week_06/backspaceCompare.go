func backspaceCompare(S string, T string) bool {
    sl, tl := len(S)-1, len(T)-1
    clearS, clearT := 0, 0

    for sl >= 0 || tl >= 0 {
        for sl >= 0 {
            if S[sl] != '#' && clearS == 0 {
                break
            }
            if S[sl] == '#' {
                clearS++
                sl--
                continue
            }
            if clearS > 0 {
                clearS--
                sl--
                continue
            }
        }

        for tl >= 0 {
            if T[tl] != '#' && clearT == 0 {
                break
            }
            if T[tl] == '#' {
                clearT++
                tl--
                continue
            }
            if clearT > 0 {
                clearT--
                tl--
                continue
            }
        }

        if sl >= 0 && tl >= 0 {
            if S[sl] != T[tl] {
                return false
            }
        } else if sl >= 0 || tl >= 0 {
            return false
        }
        sl--
        tl--
    }
    return true
}