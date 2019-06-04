package trees

// DepthFirstSearch applies a function to a tree, depth first (current node value -> left -> right)
func (t *Tree) DepthFirstSearch(fn func(*Tree) (quit bool)) {
	if t == nil {
		return
	}

	if fn != nil {
		if quit := fn(t); quit {
			return
		}
	}

	t.left.DepthFirstSearch(fn)
	t.right.DepthFirstSearch(fn)
}
