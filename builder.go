package dom

// var (
// 	walkTree func(n *node, ns []*node) []*node
// )

// Builder struct

// GetElementByClassName adds a ClassNameSelector to builder
// func (b *Builder) GetElementByClassName(s string) *Builder {
// 	b.selectors = append(b.selectors, ClassNameSelector(s))
// 	return b
// }

// // GetElementsByClassName adds a ClassNameAllSelector to builder
// func (b *Builder) GetElementsByClassName(s string) *Builder {
// 	b.selectors = append(b.selectors, ClassNameAllSelector(s))
// 	return b
// }

// // NthChild adds a NthChildSelector to builder
// func (b *Builder) NthChild(i int) *Builder {
// 	b.selectors = append(b.selectors, NthChildSelector(i))
// 	return b
// }

// // newTreeNode takes an html.Node and returns a reference to a
// // new tree node
// func newTreeNode(n *html.Node) *node {
// 	return &node{item: n}
// }

// // mapFromNode maps a slice of internal nodes to a slice of html nodes
// func mapFromNode(bs []*node) []*html.Node {
// 	bn := make([]*html.Node, len(bs))
// 	for i, n := range bs {
// 		bn[i] = n.item
// 	}

// 	return bn
// }

// // mapToNode maps a slice of html nodes to a slice of internal nodes
// func mapToNode(ns []*html.Node) []*node {
// 	bn := make([]*node, len(ns))
// 	for i, n := range ns {
// 		bn[i] = newTreeNode(n)
// 	}

// 	return bn
// }

// // Helper function that walks the current builder tree and returns a list of
// // all scalars.
// func (b *Builder) getScalars() []*node {
// 	// Create an initial slice of tNodes
// 	scalars := make([]*node, 0)

// 	// Define recursive function
// 	walkTree = func(n *node, ns []*node) []*node {
// 		// Check if scalar
// 		if len(n.children) == 0 {
// 			return append(ns, n)
// 		}

// 		for _, c := range n.children {
// 			ns = walkTree(c, ns)
// 		}

// 		return ns
// 	}

// 	return walkTree(b.root, scalars)
// }

// Build will successively execute each selector and update the
// node tree with the results
// func (b *Builder) Build(n *Node) *Builder {
// 	b.refs = []*Node{n}

// 	// Loop over selectors
// 	for _, fn := range b.selectors {
// 		items := make([]*Node, 0)

// 		// Apply selector to each scalar
// 		for _, n := range b.refs {
// 			items = append(items, fn(n)...)
// 		}

// 		b.refs = items
// 	}

// 	fmt.Println(b.refs[0])

// 	return b
// }

// // Join will collect the current tree scalars and return the html nodes
// func (b *Builder) Join() []*html.Node {
// 	return mapFromNode(b.getScalars())
// }

// NewBuilder returns a reference to a new builder instance
// func NewBuilder() *Builder {
// 	return &Builder{}
// }

// func (b *Builder) QuerySelector(s string) *Builder {
// 	qs := MakeQuerySelector(s)
// 	b.selectors = append(b.selectors, wrapSelector(qs))
// 	return b
// }

// wrapSelector wraps a selector function and transforms the single output of the
// selector to a slice
// func wrapSelector(s Selector) func(n *html.Node) []*html.Node {
// 	return func(n *html.Node) []*html.Node {
// 		val := s(n)

// 		if val == nil {
// 			return make([]*html.Node, 0)
// 		}

// 		return []*html.Node{val}
// 	}
// }
