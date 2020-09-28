package main

import "fmt"

func factorial(n int) uint32 {
	var fact uint32 = 1
	for i := 2; i <= n; i++ {
		fact *= uint32(i)
	}
	return fact
}

func main() {
	var e float64
	var max uint8

	fmt.Scanln(&max)

	for i := 0; i <= int(max); i++ {
		e += 1.0 / float64(factorial(i))
	}

	fmt.Println(e)
}
