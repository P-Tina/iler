package main

import "fmt"

func main() {
	sum := 0
	for _, i := range calculation(2000000) {
		sum += i
	}
	fmt.Println(sum)
	//пока не знаю
}
func calculation(n int) []int {
	var str []int
	for i := 0; i < n; i++ {
		str = append(str, i)
	}
	str[1] = 0
	for i := 2; i < n; i++ {
		if str[i] != 0 {
			for p := 2 * i; p < n; {
				str[p] = 0
				p += i
			}
		}
	}
	return str
}
