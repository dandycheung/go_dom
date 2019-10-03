package dom

import (
	"strings"
)

// SplitClasslist takes a *Node and splits
// the nodes class string and returns the string slice
func SplitClasslist(n *Node) []string {
	classnames := GetAttribute(n, "class")

	// Split the classnames into a slice
	return strings.Fields(classnames)
}

// BuildClasslist takes a selector string ".button.primary" and
// returns a trimmed string slice containing the classnames [button primary]
// with all trailing whitespace removed
func BuildClasslist(s string) []string {
	return Map(
		Filter(strings.Split(s, "."), IsNotEmpty),
		strings.TrimSpace,
	)
}
