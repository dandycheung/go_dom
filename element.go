package dom

import (
	"golang.org/x/net/html"
)

// GetAttribute accepts an *html.Node and a string key and returns the attribute value
func GetAttribute(n *html.Node, k string) string {
	var s string

	for _, a := range n.Attr {
		if a.Key == k {
			s = a.Val
			break
		}
	}

	return s
}

// // GetElementByClassname accepts an html.Node and classname string and
// // returns the first html node that contains the classname.
// func GetElementByClassname(c string, doc *html.Node) *html.Node {
// 	fields := strings.Fields(c)

// 	fn := func(n *html.Node) bool {
// 		return ElementContainsClass(n, fields)
// 	}

// 	return FindOne(doc, fn)
// }

// // GetElementsByClassname accepts an html.Node and classname string and
// // returns a slice of all html nodes that contains the classname.
// func GetElementsByClassname(c string, doc *html.Node) []*html.Node {
// 	fields := strings.Fields(c)

// 	fn := func(n *html.Node) bool {
// 		return ElementContainsClass(n, fields)
// 	}

// 	return FindAll(doc, fn, make([]*html.Node, 0))
// }

// // GetElementByDataType func
// func GetElementByDataType(s string, doc *html.Node) *html.Node {
// 	fn := func(n *html.Node) bool {
// 		return ElementIsDataType(n, s)
// 	}

// 	return FindOne(doc, fn)
// }

// // GetElementsByDataType func
// func GetElementsByDataType(s string, doc *html.Node) []*html.Node {
// 	fn := func(n *html.Node) bool {
// 		return ElementIsDataType(n, s)
// 	}

// 	return FindAll(doc, fn, make([]*html.Node, 0))
// }
