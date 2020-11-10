package main

import "fmt"

func xy() (int, int) {
	x := 1
	y := 2
	return x, y
}

func main() {
	x, y := xy()
	fmt.Println(xy())
	xobr(&x)
	yobr(&y)
	fmt.Println(x, y)
}
func xobr(xPtr *int) {
	*xPtr = 2
}
func yobr(yPtr *int) {
	*yPtr = 1
}
