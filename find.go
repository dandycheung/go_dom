package dom

import (
	"golang.org/x/net/html"
)

// FindOne recurses *html.Node and children returning first node where Predicate returns true
func FindOne(n *html.Node, fn Predicate) *html.Node {
	if fn(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := FindOne(c, fn)

		if found != nil {
			return found
		}
	}

	return nil
}

// FindAll recurses *html.Node returning all child *html.Node that pass the Predicate
func FindAll(n *html.Node, fn Predicate) []*html.Node {
	var walkNode func(n *html.Node, fn Predicate, ns []*html.Node) []*html.Node

	walkNode = func(n *html.Node, fn Predicate, ns []*html.Node) []*html.Node {
		if fn(n) {
			ns = append(ns, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ns = walkNode(c, fn, ns)
		}

		return ns
	}

	return walkNode(n, fn, make([]*html.Node, 0))
}
