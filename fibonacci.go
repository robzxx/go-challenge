package piscine

func Fibonacci(index int) int {
	var a int
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	b := 0
	c := 1

	for i := 2; i <= index; i++ {
		a = b + c
		b = c
		c = a
	}
	return c
}
