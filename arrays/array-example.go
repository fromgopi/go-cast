package main

import "fmt"

func main() {

	var array [5]int
	fmt.Println("Empty Array", array)

	array[4] = 100
	fmt.Println("Set Element Array: ", array)

	fmt.Println("Get ", array[4])

	fmt.Println("Length of the array: ", len(array))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
