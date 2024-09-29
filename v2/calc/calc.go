package calc

func Add(args ...int) int {

	s := 0
	for _, value := range args {
		s += value
	}
	return s
}

func Subtract(a int, b int) int {
	return (a - b)
}
