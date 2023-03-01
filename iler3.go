package main

import "fmt"

func main() {
	n := 2
	x := 600851475143
	for i := 0; i <= x; i++ {
		if x%n == 0 {
			x = x / n
		} else {
			n += 1
		}
	}
	fmt.Println(n)
}
