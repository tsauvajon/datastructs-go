package trees

import (
	"testing"
)

func TestDFS(t *testing.T) {
	demoTree := createTree()

	output := ""

	demoTree.DepthFirstSearch(func(currentTree *Tree) (quit bool) {
		output = output + currentTree.value
		return false
	})

	expected := "ABCDEFGHIJKLMNO"

	if output != expected {
		t.Errorf("expected %s, got %s\n", expected, output)
	}
}

// Tess that BFS can find a value
func TestFindingDFS(t *testing.T) {
	demoTree := createTree()

	var (
		lf    string
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

	if err := demoTree.DepthFirstSearch(find(lf)); err != nil {
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
		t.Errorf("should not have found %s but it did", lf)
	}
}
