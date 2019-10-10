package dom

import "golang.org/x/net/html"

// IterateSiblings iterates through an html.Node and all her siblings, calling
// func arg for each iteration
func IterateSiblings(n *html.Node, fn func(n *html.Node)) {
	for s := n; s != nil; s = s.NextSibling {
		fn(s)
	}
}
