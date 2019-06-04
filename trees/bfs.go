package trees

import "errors"

func breadthFirstSearch(t *Tree, lookingFor string) (*Tree, error) {
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

		if node.value == lookingFor {
			return node, nil
		}
	}

	return nil, errors.New("not found")
}
