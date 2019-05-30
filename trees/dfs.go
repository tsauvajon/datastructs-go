package main

import (
	"errors"
)

func depthFirstSearch(t *tree, lookingFor string) (*tree, error) {
	if t == nil {
		return nil, errors.New("end of tree reached")
	}

	if n, _ := depthFirstSearch(t.left, lookingFor); n != nil {
		return n, nil
	} else if n, _ := depthFirstSearch(t.right, lookingFor); n != nil {
		return n, nil
	}

	if t.value == lookingFor {
		return t, nil
	}

	return nil, errors.New("not found")
}
