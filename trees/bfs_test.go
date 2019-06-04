package trees

import (
	"testing"
)

func TestBFS(t *testing.T) {
	demoTree := createTree()

	output := ""

	if err := BreadthFirstSearch(demoTree, func(currentTree *Tree) {
		output = output + currentTree.value
	}); err != nil {
		t.Errorf("the function failed: %v", err)
	}

	expected := "ABICFJMDEGHKLNO"

	if output != expected {
		t.Errorf("expected %s, got %s\n", expected, output)
	}
}
