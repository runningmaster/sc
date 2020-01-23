package parser

import (
	"strconv"
)

type Node struct {
	typ   TokenType
	prev  *Node
	next  []*Node
	vals  []string
	depth int
}

func (n *Node) Type() TokenType {
	return n.typ
}

func (n *Node) Vals() []string {
	return n.vals
}

func (n *Node) Depth() int {
	return n.depth
}

func (n *Node) String() string {
	return strconv.Itoa(n.depth) + " ->" +
		" type:" + n.typ.String() +
		" next:" + strconv.Itoa(len(n.next)) +
		" vals:" + strconv.Itoa(len(n.vals))
}

type WalkFunc func(node *Node, err error) error

// Walk traverses from Root.
func (n *Node) Walk(walkFn WalkFunc) error {
	err := walkFn(n, nil)
	if err != nil {
		return err
	}

	for _, n = range n.next {
		err = n.Walk(walkFn)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply traverses from Left-SubTree, Right-SubTree, and Root.
func (n *Node) Apply(walkFn WalkFunc) error {
	return apply(n, walkFn)
}

func apply(n *Node, walkFn WalkFunc) error {
	if n == nil {
		return nil
	}

	var err error
	for _, v := range n.next {
		err = apply(v, walkFn)
		if err != nil {
			return err
		}
	}

	return walkFn(n, err)
}
