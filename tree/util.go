package tree

import "fmt"

func (root *node) traverseInorder() {
	curr := root
	for curr != nil {
		curr.left.traverseInorder()
		fmt.Printf(" %d ", curr.val)
		curr.right.traverseInorder()
		break
	}
}

func (root *node) traversePreorder() {
	curr := root
	for curr != nil {
		fmt.Printf(" %d ", curr.val)
		curr.left.traversePreorder()
		curr.right.traversePreorder()
		break
	}
}

func (root *node) traversePostorder() {
	curr := root
	for curr != nil {
		curr.left.traversePostorder()
		curr.right.traversePostorder()
		fmt.Printf(" %d ", curr.val)
		break
	}
}

func (root *node) height() int {

	if root == nil {
		return 0
	}

	// call Height for left and right subtree
	lheight := root.left.height()
	rheight := root.right.height()

	if lheight > rheight {
		return lheight + 1
	}

	return rheight + 1
}
