func replaceSpace(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch == ' ' {
			sb.WriteString("%20")
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}