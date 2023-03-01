package main

import (
	"fmt"
)

func main() {
	sum := 0
	a := 0
	for i := 0; i <= 100; i++ {
		sum = sum + i
		a = a + i*i
	}
	kv := sum*sum - a

	fmt.Println(kv, sum, a)

}
