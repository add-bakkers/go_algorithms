package main

import "fmt"

type SplayTree struct {
	root *Node
}

type Node struct {
	value  interface{}
	parent *Node
	left   *Node
	right  *Node
}

func rotate(child *Node) error {
	var parent *Node = child.parent
	if parent == nil {
		return fmt.Errorf("parent is nil")
	}
	if parent.left == child {
		parent.left = child.right
		if child.right != nil {
			child.right.parent = parent
		}
		child.right = parent
		return nil
	} else {
		parent.right = child.left
		if child.left != nil {
			child.left.parent = parent
		}
		child.left = parent
		return nil
	}
}

func zig(n *Node) {
	rotate(n)
}

func zigzig(child, parent *Node) {
	rotate(parent)
	rotate(child)
}

func zigzag(n *Node) {
	rotate(n)
	rotate(n)
}

func (t *SplayTree) Splay(n *Node) {
	for n.parent != nil {
		var parent = n.parent
		if parent.parent == nil {
			zig(n)
		} else {
			if n == parent.left && parent == parent.parent.left {
				zigzig(n, parent)
			} else {
				zigzag(n)
			}
		}
	}
	t.root = n
}
