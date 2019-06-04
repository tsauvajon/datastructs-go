package trees

// DepthFirstSearch applies a function to a tree, depth first (current node value -> left -> right)
func DepthFirstSearch(t *Tree, fn func(*Tree)) {
	if t == nil {
		return
	}

	if fn != nil {
		fn(t)
	}

	DepthFirstSearch(t.left, fn)
	DepthFirstSearch(t.right, fn)
}
