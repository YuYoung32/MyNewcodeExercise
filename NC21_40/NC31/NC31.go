package NC31

func FirstNotRepeatingChar(str string) int {
	m := make([]byte, 128)
	ret := -1
	for _, char := range str {
		m[char]++
	}
	for i, char := range str {
		if m[char] == 1 {
			ret = i
			break
		}
	}
	return ret
}
