package main

func main() {
	i := 2
	count := 0
	for i := 1; i > 0; i++ {
		if chet(i) {
			count++
			if count == 10001 {
				break
			}
		}
	}
}

func chet(int a) bool {
	for i := 2; i < a; i++ {
		{
			if a%i == 0 {
				return false
			}
		}
	}
	return true
	//----------------
}
