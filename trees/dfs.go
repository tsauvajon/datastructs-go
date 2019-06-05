package trees

import "errors"

// DepthFirstSearch applies a function to a tree, depth first
// (current node value -> left -> right)
func (t *Tree) DepthFirstSearch(fn func(*Tree) (quit bool)) error {
	if t == nil {
		return errors.New("t cannot be nil")
	}

	if fn != nil {
		if quit := fn(t); quit {
			return nil
		}
	}

	t.left.DepthFirstSearch(fn)
	t.right.DepthFirstSearch(fn)

	return nil
}
