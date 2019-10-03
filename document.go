package dom

import "golang.org/x/net/html"

// GetDocumentBody recurses *html.Node until a <body> node is found or returns nil
func GetDocumentBody(doc *html.Node) *html.Node {
	pred := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "body"
	}

	var find func(n *html.Node, fn func(n *html.Node) bool) *html.Node

	find = func(n *html.Node, fn func(n *html.Node) bool) *html.Node {
		if fn(n) {
			return n
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			found := find(c, fn)

			if found != nil {
				return found
			}
		}

		return nil
	}

	return find(doc, pred)
}

// // FindAll recurses *html.Node returning all child *html.Node that pass the DomPredicate
// func FindAll(n *html.Node, fn DomPredicate, nodes []*html.Node) []*html.Node {
// 	if fn(n) {
// 		nodes = append(nodes, n)
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		nodes = FindAll(c, fn, nodes)
// 	}

// 	return nodes
// }

// // ElementIsDataType accepts an html.Node and a string data type
// // and returns true if node matches data type
// func ElementIsDataType(n *html.Node, s string) bool {
// 	return n.Data == s
// }

// // ElementContainsClass accepts an html.Node and a string of classnames
// // and returns a bool determining if the node classlist attribute contains each class
// func ElementContainsClass(n *html.Node, fields []string) bool {
// 	cx := SplitNodeClasslist(n)
// 	counter := 0

// 	for _, x := range fields {
// 		if ContainsString(cx, x) {
// 			counter++
// 		}
// 	}

// 	return len(fields) == counter
// }

// // ParseElementID returns id string attr of given *html.Node
// func ParseElementID(n *html.Node) string {
// 	var id string

// 	for _, a := range n.Attr {
// 		if a.Key == "id" {
// 			id = a.Val
// 			break
// 		}
// 	}

// 	return id
// }

// // ParseNodeID returns id attr of node
// func ParseNodeID(n *html.Node) string {
// 	var id string

// 	for _, a := range n.Attr {
// 		if a.Key == "id" {
// 			id = a.Val
// 			break
// 		}
// 	}

// 	return id
// }
