package trees

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