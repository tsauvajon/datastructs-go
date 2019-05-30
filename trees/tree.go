package main

type tree struct {
	value string
	left  *tree
	right *tree
}

func createTree() *tree {
	t := tree{
		value: "A",
		left: &tree{
			value: "B",
			left: &tree{
				value: "C",
				left: &tree{
					value: "D",
				},
				right: &tree{
					value: "E",
				},
			},
			right: &tree{
				value: "F",
				left: &tree{
					value: "G",
				},
				right: &tree{
					value: "H",
				},
			},
		},
		right: &tree{
			value: "I",
			left: &tree{
				value: "J",
				left: &tree{
					value: "K",
				},
				right: &tree{
					value: "L",
				},
			},
			right: &tree{
				value: "M",
				left: &tree{
					value: "N",
				},
				right: &tree{
					value: "O",
				},
			},
		},
	}

	return &t
}
