package main

import "fmt"

func main() {
	var a int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			a += i
		}
	}
	fmt.Println(a)
}
