package trees

import (
	"testing"
)

func TestDFS(t *testing.T) {
	demoTree := createTree()

	output := ""

	demoTree.DepthFirstSearch(func(currentTree *Tree) {
		output = output + currentTree.value
	})

	expected := "ABCDEFGHIJKLMNO"

	if output != expected {
		t.Errorf("expected %s, got %s\n", expected, output)
	}
}
