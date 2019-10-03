package dom

var walkNode func(n *Node, fn Predicate, ns []*Node) []*Node

// FindOne recurses *Node and children returning first node where Predicate returns true
func FindOne(n *Node, fn Predicate) *Node {
	if fn(n) {
		return n
	}

	for i := range n.Children {
		found := FindOne(n.Children[i], fn)

		if found != nil {
			return found
		}
	}

	return nil
}

// FindAll recurses *Node returning all child *Node that pass the Predicate
func FindAll(n *Node, fn Predicate) []*Node {
	walkNode = func(n *Node, fn Predicate, ns []*Node) []*Node {
		if fn(n) {
			ns = append(ns, n)
		}

		for i := range n.Children {
			ns = walkNode(n.Children[i], fn, ns)
		}

		return ns
	}

	return walkNode(n, fn, make([]*Node, 0))
}
