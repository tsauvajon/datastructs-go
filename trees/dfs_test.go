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
