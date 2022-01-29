package MyTools

func Compare2DArray(a1 [][]int, a2 [][]int) bool {
	compare1DArray := func(a1, a2 []int) bool {
		if len(a1) != len(a2) {
			return false
		}
		for i := 0; i < len(a1); i++ {
			if a1[i] != a2[i] {
				return false
			}
		}
		return true
	}

	if len(a1) != len(a2) {
		return false
	}

	flags := make([]bool, len(a2))
	for i := 0; i < len(flags); i++ {
		flags[i] = false
	}

	for i1 := 0; i1 < len(a1); i1++ {
		for i2 := 0; i2 < len(a2); i2++ {
			if flags[i2] == false && compare1DArray(a1[i1], a2[i2]) {
				flags[i2] = true
				break
			}
		}
	}
	for _, f := range flags {
		if f == false {
			return false
		}
	}
	return true
}
