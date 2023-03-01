package main

import "fmt"

func main() {
	var a, sum int
	x := 0
	y := 1
	for i := 1; i > 0; i++ {
		a = x + y
		x = y
		y = a
		if a%2 == 0 {
			sum += a
		}
		if a >= 4000000 {
			break
		}
	}
	fmt.Println(sum)
}
