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
func TagSelectorAll(t string) SelectorAll {
	predicate := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == t
	}

	return func(n *html.Node) []*html.Node {
		return FindAll(n, predicate)
	}
}

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

// ClassNameSelectorAll returns a selector that finds all the nodes
// containing every item in classlist.
func ClassNameSelectorAll(s string) SelectorAll {
	classlist := BuildClasslist(s)

	predicate := func(n *html.Node) bool {
		return ContainsClassname(n, classlist)
	}

	return func(n *html.Node) []*html.Node {
		return FindAll(n, predicate)
	}
}

// MakeQuerySelectorAll is a higher order function that takes a selector query string,
// parses the query string into a slice of selectors, and returns a SelectorAll
func MakeQuerySelectorAll(q string) SelectorAll {
	qs := strings.Fields(q)
	fns := make([]SelectorAll, 0)

	for _, s := range qs {
		fns = append(fns, parseSelectorQueryAll(s))
	}

	return func(n *html.Node) []*html.Node {
		ns := []*html.Node{n}

		for _, fn := range fns {
			fs := make([]*html.Node, 0)

			for _, node := range ns {
				fs = append(fs, fn(node)...)
			}

			if len(fs) == 0 {
				break
			}

			ns = fs
		}

		return ns
	}
}

// MakeQuerySelector is a higher order function that takes a selector query string,
// parses the query string into a slice of selectors, and returns a Selector
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

func parseSelectorQueryAll(q string) SelectorAll {
	firstCh := q[0]

	switch string(firstCh) {
	case ".":
		return ClassNameSelectorAll(q)
	default:
		return TagSelectorAll(q)
	}
}
