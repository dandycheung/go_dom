package dom

import (
	"strings"

	"golang.org/x/net/html"
)

// Selector type represents a function that accepts an html node
// and returns a single html node
type Selector func(n *html.Node) *html.Node

// SelectorAll type represents a function that accepts an html node
// and returns a slice of html nodes
type SelectorAll func(n *html.Node) []*html.Node

// NthChildSelector returns a selector that returns the child node at given a index
// func NthChildSelector(i int) Selector {
// 	return func(n *html.Node) []*html.Node {
// 		f := n.FirstChild

// 		for x := 0; x <= i; x++ {
// 			if x == i {
// 				return []*html.Node{f}
// 			}
// 			f = f.NextSibling
// 		}

// 		return nil
// 	}
// }

// TagSelector returns a selector that returns this first node
// that matches the provided tag
func TagSelector(t string) Selector {
	predicate := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == t
	}

	return func(n *html.Node) *html.Node {
		return FindOne(n, predicate)
	}
}

// TagSelectorAll returns a selector that finds all nodes matching given tag name
// func TagSelectorAll(t string) Selector {
// 	predicate := func(n *html.Node) bool {
// 		return n.Type == html.ElementNode && n.Data == t
// 	}

// 	return func(n *html.Node) []*html.Node {
// 		return FindAll(n, predicate)
// 	}
// }

// ClassNameSelector returns a selector that finds the first node
// containing every item in classlist.
func ClassNameSelector(s string) Selector {
	classlist := BuildClasslist(s)

	predicate := func(n *html.Node) bool {
		return ContainsClassname(n, classlist)
	}

	return func(n *html.Node) *html.Node {
		return FindOne(n, predicate)
	}
}

// ClassNameAllSelector returns a selector that finds all the nodes
// containing every item in classlist.
// func ClassNameAllSelector(s string) Selector {
// 	classlist := BuildClasslist(s)

// 	predicate := func(n *html.Node) bool {
// 		return ContainsClassname(n, classlist)
// 	}

// 	return func(n *html.Node) []*html.Node {
// 		return FindAll(n, predicate)
// 	}
// }

// MakeQuerySelector is a higher order function that takes a selector query string,
// parses the query string into a slice of selectors, and returns a wrapped selector function
func MakeQuerySelector(q string) Selector {
	qs := strings.Fields(q)
	fns := make([]Selector, 0)

	for _, s := range qs {
		fns = append(fns, parseSelectorQuery(s))
	}

	return func(n *html.Node) *html.Node {
		for _, fn := range fns {
			n = fn(n)
			if n == nil {
				break
			}
		}

		return n
	}
}

func parseSelectorQuery(q string) Selector {
	firstCh := q[0]

	switch string(firstCh) {
	case ".":
		return ClassNameSelector(q)
	default:
		return TagSelector(q)
	}
}
