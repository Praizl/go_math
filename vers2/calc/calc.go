package calc

func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v

	}
	return s
}

func Subtract(a int, b int) int {
	return (a - b)
}
