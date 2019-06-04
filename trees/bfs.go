package trees

import "errors"

// BreadthFirstSearch applies a function to a tree, breadth first (current node value -> all of its children -> all of their children)
func BreadthFirstSearch(t *Tree, fn func(*Tree)) error {
	if t == nil {
		return errors.New("t cannot be nil")
	}

	q := []*Tree{t}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.left != nil {
			q = append(q, node.left)
		}
		if node.right != nil {
			q = append(q, node.right)
		}

		fn(node)
	}

	return nil
}
