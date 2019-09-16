package dom

// ContainsString returns true if string is an element of []string
func ContainsString(cx []string, s string) bool {
	for _, x := range cx {
		if x == s {
			return true
		}
	}
	return false
}
