package main

import "fmt"

func main() {
	x := 20
	var str []int
	for j := 2; j < x; j++ {
		f := true
		for i := 2; i < j; i++ {
			if j%i == 0 {
				f = false
			}
		}
		if f == true {
			str = append(str, j)
		}
	}
	fmt.Println(str)
	//пока не знаю
}
