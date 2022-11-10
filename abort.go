package piscine

func Abort(a, b, c, d, e int) int {
	middle := SortInTable(a, b, c, d, e)[2]
	return middle
}

func SortInTable(a, b, c, d, e int) []int {
	emma := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if emma[i] > emma[j] {
				emma[i], emma[j] = emma[j], emma[i]
			}
		}
	}
	return emma
}
