package MyTools

//Compare2DArray 比较二维数组，忽略在二维上的顺序
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

//StrictCompare2DArray 比较二维数组，同时注意在二维上的顺序
func StrictCompare2DArray(a1 [][]int, a2 [][]int) bool {
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

	for i := 0; i < len(a1); i++ {
		if compare1DArray(a1[i], a2[i]) == false {
			return false
		}
	}
	return true
}

//CompareStringArray 比较字符串数组，忽略在二维上的顺序
func CompareStringArray(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	flags := make([]bool, len(s1))
	for i := 0; i < len(flags); i++ {
		flags[i] = false
	}
	for _, s := range s1 {
		for i, cs := range s2 {
			if flags[i] == false && s == cs {
				flags[i] = true
			}
		}
	}
	for _, flag := range flags {
		if flag == false {
			return false
		}
	}
	return true
}

//Compare1DArray 比较一维数组是否相等
func Compare1DArray(a1, a2 []int) bool {
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
