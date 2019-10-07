package dom

import "golang.org/x/net/html"

// GetAttribute accepts an *html.Node and a string key and returns the attribute value
func GetAttribute(n *html.Node, k string) string {
	var s string

	for _, a := range n.Attr {
		if a.Key == k {
			s = a.Val
			break
		}
	}

	return s
}
