package dom

import "golang.org/x/net/html"

// Predicate func type takes an *html.Node and returns bool
type Predicate func(n *html.Node) bool
