package dom

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

var walkTree func(root *html.Node) *Node

type Node struct {
	Tag        string
	ID         string
	Data       interface{}
	Children   []*Node
	ClassNames []string
	Type       html.NodeType
}

func makeNode(n *html.Node) *Node {
	return &Node{
		Tag:        n.Data,
		ClassNames: BuildClasslist(GetAttribute(n, "class")),
		ID:         GetAttribute(n, "id"),
		Type:       n.Type,
	}
}

func ParseHTMLNode(root *html.Node) *Node {
	walkTree := func(root *html.Node) *Node {
		fmt.Println(root.Type)
		node := makeNode(root)

		for c := root.FirstChild; c != nil; c = c.NextSibling {
			node.Children = append(node.Children, walkTree(c))
		}

		return node
	}

	return walkTree(root)
}

func ParseBody(r io.ReadCloser) (*Node, error) {
	defer r.Close()

	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	return ParseHTMLNode(GetDocumentBody(doc)), nil
}

func (n *Node) GetElementByClassname(cs string) *Node {
	classlist := BuildClasslist(cs)

	predicate := func(n *Node) bool {
		return ContainsClassname(n, classlist)
	}

	return FindOne(n, predicate)
}
