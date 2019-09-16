package dom

type Node struct {
	tag        string
	data       interface{}
	children   []*Node
	classNames []string
}
