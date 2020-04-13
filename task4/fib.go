package task4

func fibonacci() func() int {
	curr, next := 0, 1
	return func() int {
		curr, next = next, curr + next
		return curr
	}
}
