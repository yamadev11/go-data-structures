package tree

import "fmt"

func NewBinarySearchTree(values ...int) *BinarySearchTree {

	var bst *BinarySearchTree
	for i, value := range values {
		if i == 0 {
			bst.root = &node{val: value}
			continue
		}

		bst.AddNode(value)
	}

	return bst
}

func (bst *BinarySearchTree) AddNode(value int) {
	curr := bst.root
	for curr != nil {
		if curr.val == value {
			fmt.Println("Node is already present", curr.val, value)
			return
		}

		if value < curr.val {
			if curr.left == nil {
				fmt.Println(value, "is added to left of", curr.val)
				curr.left = &node{val: value}
				return
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				fmt.Println(value, "is added to right of", curr.val)
				curr.right = &node{val: value}
				return
			}
			curr = curr.right
		}
	}
}

func (bst *BinarySearchTree) TraverseInorder() {
	if bst == nil {
		// log error
		return
	}

	bst.root.traverseInorder()
}

func (bst *BinarySearchTree) TraversePreorder() {
	if bst == nil {
		// log error
		return
	}

	bst.root.traversePreorder()
}

func (bst *BinarySearchTree) TraversePostorder() {
	if bst == nil {
		// log error
		return
	}

	bst.root.traversePostorder()
}

func (bst *BinarySearchTree) Height() int {
	if bst == nil {
		// log error
		return 0
	}

	return bst.root.height()
}

func (root *node) LeftView() {
	// need to correct the logic
	curr := root
	for curr != nil {
		fmt.Printf(" %d ", curr.val)
		if curr.left != nil {
			curr = curr.left
			continue
		}
		curr = curr.right
	}
}

func (root *node) RightView() {
	// need to correct the logic
	curr := root
	for curr != nil {
		fmt.Printf(" %d ", curr.val)
		if curr.right != nil {
			curr = curr.right
			continue
		}
		curr = curr.left
	}
}
