package dom

import "golang.org/x/net/html"

// GetChildByType returns a the first child node that have same type as NodeType
func GetChildByType(n *html.Node, t html.NodeType) (f *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == t {
			f = c
			break
		}
	}

	return f
}

// GetChildrenByType returns a list of child nodes that have same type as NodeType
func GetChildrenByType(n *html.Node, t html.NodeType) []*html.Node {
	ns := make([]*html.Node, 0)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == t {
			ns = append(ns, c)
		}
	}

	return ns
}
