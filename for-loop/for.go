package main

import "fmt"

func main() {

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i *= 2
	}
	fmt.Println()
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("Prime numbers ")

	for n := 1; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
