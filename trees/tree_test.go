package trees

import "testing"

func createTree() *Tree {
	t := &Tree{
		value: "A",
		left: &Tree{
			value: "B",
			left: &Tree{
				value: "C",
				left: &Tree{
					value: "D",
				},
				right: &Tree{
					value: "E",
				},
			},
			right: &Tree{
				value: "F",
				left: &Tree{
					value: "G",
				},
				right: &Tree{
					value: "H",
				},
			},
		},
		right: &Tree{
			value: "I",
			left: &Tree{
				value: "J",
				left: &Tree{
					value: "K",
				},
				right: &Tree{
					value: "L",
				},
			},
			right: &Tree{
				value: "M",
				left: &Tree{
					value: "N",
				},
				right: &Tree{
					value: "O",
				},
			},
		},
	}

	return t
}

func TestTree (t *testing.T) {
	tree := createTree()

	actualLeft := tree.left.left.left
	actualRight := tree.right.right.right

	expectedLeftVal := "D"
	expectedRightVal := "O"

	if actualLeft.value != expectedLeftVal {
		t.Errorf("expected leftmost node to be %s, got %s", expectedLeftVal, actualLeft.value)
	}

	if actualRight.value != expectedRightVal {
		t.Errorf("expected rightmost node to be %s, got %s", expectedRightVal, actualRight.value)
	}
}
