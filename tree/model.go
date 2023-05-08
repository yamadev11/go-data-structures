package tree

type node struct {
	val   int
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

type BinarySearchTree struct {
	root *node
}
