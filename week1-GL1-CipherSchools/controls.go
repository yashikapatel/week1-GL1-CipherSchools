package main

import "fmt"

func main() {

	data :=1009
	switch data {
		case10:
		data = 1000
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough
	case 15:
		data = 1002
		fmt.Println(data)
	default:
		data = 10909
		fmt.Println(data)
	}
}

	fmt.Print("Enter a Number:")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("hey you are even")
	} else {
		fmt.Println("hey you are odd")
	}

	x := 10000
	if x:=10; x%2 == 0 {
		fmt.Println("hey you are even")
		x = 200
	} else {
		fmt.Println("hey you are odd")
	}
	fmt.Println(x)
}
