package dom

import "golang.org/x/net/html"

// ContainsClassname accepts a Node and a string of classnames
// and returns a bool determining if the node classlist attribute contains each class
func ContainsClassname(n *html.Node, fields []string) bool {
	cx := SplitClasslist(n)
	counter := 0

	for _, x := range fields {
		if ContainsString(cx, x) {
			counter++
		}
	}

	return len(fields) == counter
}
