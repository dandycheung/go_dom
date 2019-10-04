package dom

import (
	"golang.org/x/net/html"
)

// QuerySelector returns the first element that matches a specified selector in the document
func QuerySelector(s string, n *html.Node) *html.Node {
	return MakeQuerySelector(s)(n)
}

// QuerySelectorAll returns all elements that match a specified selector in the document
func QuerySelectorAll(s string, n *html.Node) []*html.Node {
	return MakeQuerySelectorAll(s)(n)
}
