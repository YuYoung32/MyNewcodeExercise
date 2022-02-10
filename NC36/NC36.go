package NC36

func findMedianinTwoSortedAray(arr1 []int, arr2 []int) int {
	// write code here
	target := (len(arr1) + len(arr2)) / 2
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			i++
			if i+j == target {
				return arr1[i-1]
			}
		} else {
			j++
			if i+j == target {
				return arr2[j-1]
			}
		}
	}
	return 0
}
