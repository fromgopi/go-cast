package main

import (
	"fmt"
	"math"
)

const str string = "This is Constant string"

func main() {
	fmt.Println(math.Abs(39.35))
	fmt.Println(str)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
