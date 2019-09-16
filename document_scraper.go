package dom

/*
package queryselector

import (
	"fmt"

	"golang.org/x/net/html"
)

type DocumentScraper struct {
	root    *node
	depth   int
	filters []FilterPredicate
}

type FilterResult struct {
	nodes []*html.Node
}

type node struct {
	tag      *html.Node
	depth    int

}

type FilterPredicate func(*html.Node) *FilterResult

func (s *DocumentScraper) getScalars() []*node {
	scalars := make([]*node, 0)

	checkDepth := func(n *node) {
		if n.depth == s.depth {
			scalars = append(scalars, n)
		}
	}

	applyTraverseTree(s.root, checkDepth)

	return scalars
}

func applyTraverseTree(root *node, fn func(*node)) {
	fn(root)

	for _, n := range root.children {
		applyTraverseTree(n, fn)
	}
}

// func (s *DocumentScraper) FilterClassname(cx string) *DocumentScraper {
// 	filter := func(doc *html.Node) *html.Node {
// 		return GetElementByClassname(doc, cx)
// 	}

// 	return s.applyFilter(key, filter)
// }

func (s *DocumentScraper) applyFilter(filter FilterPredicate) *DocumentScraper {
	for _, scalar := range s.getScalars() {

		result := filter(scalar.tag)

		for _, n := range result.nodes {
			scalar.children = append(scalar.children, makeNode(n, s.depth+1))
		}
	}

	s.updateDepth()

	return s
}

func (s *DocumentScraper) updateDepth() {
	checkDepth := func(n *node) {
		if n.depth > s.depth {
			s.depth = n.depth
		}
	}

	applyTraverseTree(s.root, checkDepth)
}

func (s *DocumentScraper) FilterTable() *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		n := GetElementByDataType(doc, "table")

		if n != nil {
			result.nodes = append(result.nodes, n)
		}

		return result
	}

	s.filters = append(s.filters, filter)
	return s
}

func (s *DocumentScraper) MFilterTable() *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		result.nodes = append(result.nodes, GetElementByDataType(doc, "table"))

		return result
	}
	s.filters = append(s.filters, filter)
	return s
}

func (s *DocumentScraper) FilterTableBody() *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		n := GetElementByDataType(doc, "tbody")

		if n != nil {
			result.nodes = append(result.nodes, n)
		}

		return result
	}

	s.filters = append(s.filters, filter)
	return s
}

func (s *DocumentScraper) FilterTableRow() *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		n := GetElementByDataType(doc, "tr")

		if n != nil {
			result.nodes = append(result.nodes, n)
		}

		return result
	}

	s.filters = append(s.filters, filter)
	return s
}

func (s *DocumentScraper) MFilterTableRow() *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		result.nodes = append(result.nodes, GetElementsByDataType(doc, "tr")...)

		return result
	}

	s.filters = append(s.filters, filter)
	return s
}

func (s *DocumentScraper) NthChild(nth int) *DocumentScraper {
	filter := func(doc *html.Node) *FilterResult {
		idx := 1

		result := &FilterResult{
			nodes: make([]*html.Node, 0),
		}

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			if idx == nth {
				result.nodes = append(result.nodes, c)
				break
			}
			idx++
			if idx > nth {
				break
			}
		}
		return result
	}

	s.filters = append(s.filters, filter)
	return s
}

func makeNode(n *html.Node, depth int) *node {
	return &node{
		tag:      n,
		children: make([]*node, 0),
		depth:    depth,
	}
}

func (s *DocumentScraper) Map() {

}

func (s *DocumentScraper) Build(doc *html.Node) interface{} {
	s.root = makeNode(doc, 1)
	s.depth = s.root.depth

	for _, filter := range s.filters {
		s.applyFilter(filter)
	}

	item := s.getScalars()[0].tag
	fmt.Println(item)
	return nil
}

func NewDocumentScraper() *DocumentScraper {
	return &DocumentScraper{
		depth:   0,
		root:    nil,
		filters: make([]FilterPredicate, 0),
	}
}
*/
