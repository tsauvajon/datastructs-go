package main

import "fmt"

func main() {
	t := createTree()

	if n, err := breadthFirstSearch(t, "F"); err != nil {
		fmt.Printf("Didn't find node: %v\n", err)
	} else {
		fmt.Printf("Found node: %v\n", n)
	}

	if n, err := depthFirstSearch(t, "F"); err != nil {
		fmt.Printf("Didn't find node: %v\n", err)
	} else {
		fmt.Printf("Found node: %v\n", n)
	}
}
