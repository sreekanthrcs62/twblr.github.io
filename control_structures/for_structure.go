package control_structures

func collatzChainLength(i int) int {
	var count int
	for i !=1 {
		if i % 2 == 0 {
			i = i/2
		} else {
			i = 3 * i + 1
		}
		count++
	}
	return count
}
