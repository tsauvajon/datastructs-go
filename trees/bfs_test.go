package trees

import (
	"testing"
)

// Tests that BFS browses the tree in the right order
func TestBrowsingBFS(t *testing.T) {
	demoTree := createTree()

	output := ""

	if err := demoTree.BreadthFirstSearch(func(currentTree *Tree) (quit bool) {
		output = output + currentTree.value
		return false
	}); err != nil {
		t.Errorf("the function failed: %v", err)
	}

	expected := "ABICFJMDEGHKLNO"

	if output != expected {
		t.Errorf("expected %s, got %s", expected, output)
	}
}

// Tess that BFS can find a value
func TestFindingBFS(t *testing.T) {
	demoTree := createTree()

	var (
		lf string
		found bool
	)

	find := func(lookingFor string) func(*Tree) bool {
		return func(currentTree *Tree) (quit bool) {
			if lookingFor == currentTree.value {
				found = true
				return true
			}
			return false
		}
	}	

	lf = "F"
	found = false

	if err := demoTree.BreadthFirstSearch(find(lf)); err != nil {
		t.Errorf("the function failed: %v", err)
	}

	if !found {
		t.Errorf("should have found %s and did not", lf)
	}

	lf = "Z"
	found = false

	if err := demoTree.BreadthFirstSearch(find(lf)); err != nil {
		t.Errorf("the function failed: %v", err)
	}

	if found {
		t.Errorf("should have found %s and did not", lf)
	}
}
