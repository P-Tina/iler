package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b int
	for i := 900; i < 1000; i++ {
		for k := 900; k < 1000; k++ {
			c := strconv.Itoa(k * i)
			a := []rune(c)
			if a[0] == a[5] && a[1] == a[4] && a[2] == a[3] {
				b = k * i
			}
		}
	}
	fmt.Println(b)
}
