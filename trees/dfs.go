package trees

// DepthFirstSearch applies a function to a tree, depth first (current node value -> left -> right)
func (t *Tree) DepthFirstSearch(fn func(*Tree)) {
	if t == nil {
		return
	}

	if fn != nil {
		fn(t)
	}

	t.left.DepthFirstSearch(fn)
	t.right.DepthFirstSearch(fn)
}
