package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		a := i
		for j := 0; j < 1000; j++ {
			b := j
			c := 1000 - b - a
			if a < b && b < c && a*a+b*b == c*c {
				if a+b+c == 1000 {
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
