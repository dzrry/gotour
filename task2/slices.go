package task2

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			p[y][x] = uint8((x + y) / 2)
		}
	}
	return p
}
