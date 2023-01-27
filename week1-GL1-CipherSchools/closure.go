package main

import "fmt"

func main() {
	// number := 10
	var number int
	number = 10
	fmt.Println(number)

	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}

	g(getInt, 8)
	/* getInt(1)
	   getInt = func(x int) int{
		fmt.Println("In a 2nd function")
		return 10+ x
	 }*/
	var y func()
	y = func() {
		fmt.Print(9)
	}
	y()
	j := func(x int) int {
		fmt.Println("In a 2nd function")
		return 20 + x
	}(10)
	fmt.Println(j)
}
func g(getInt func(int) int, u int) {
	j := getInt(78)
	fmt.Print(j)
}
