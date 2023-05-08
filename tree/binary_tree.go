package tree

import "fmt"

func NewBinaryTree(values ...int) *BinaryTree {

	var bTree *BinaryTree
	for i, value := range values {
		if i == 0 {
			bTree.root = &node{val: value}
			continue
		}

		bTree.AddNode(value)
	}

	return bTree
}

func (bTree *BinaryTree) AddNode(value int) {
	curr := bTree.root
	var nodes []*node
	var i int
	for curr != nil {
		if curr.val == value {
			fmt.Println("Node is already present", curr.val, value)
			return
		}

		if curr.left == nil {
			fmt.Println(value, "is added to left of", curr.val)
			curr.left = &node{val: value}
			return
		}

		if curr.right == nil {
			fmt.Println(value, "is added to right of", curr.val)
			curr.right = &node{val: value}
			return
		}

		nodes = append(nodes, curr.left, curr.right)
		curr = nodes[i]
		i++
	}
}

func (bTree *BinaryTree) LevelOrder() {
	curr := bTree.root
	var nodes []*node
	var i int
	for curr != nil {

		fmt.Printf(" %d ", curr.val)
		if curr.left != nil {
			nodes = append(nodes, curr.left)
		}

		if curr.right != nil {
			nodes = append(nodes, curr.right)
		}

		if i == len(nodes) {
			return
		}

		curr = nodes[i]
		i++
	}
}

func (bTree *BinaryTree) TraverseInorder() {
	if bTree == nil {
		// log error
		return
	}

	bTree.root.traverseInorder()
}

func (bTree *BinaryTree) TraversePreorder() {
	if bTree == nil {
		// log error
		return
	}

	bTree.root.traversePreorder()
}

func (bTree *BinaryTree) TraversePostorder() {
	if bTree == nil {
		// log error
		return
	}

	bTree.root.traversePostorder()
}

func (bTree *BinaryTree) Height() int {
	if bTree == nil {
		// log error
		return 0
	}

	return bTree.root.height()
}
