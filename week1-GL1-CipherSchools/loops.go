package main

import "fmt"

func main() {
	/*for i := 0; i < 3; i++ {
	for i := 0; i < 3; i++
	fmt.Println(i)*/
	//}
	//}
	// for true{
	// 	fmt.Println("INfinite Execution")
	// }
	// i:=0
	// for i<3{
	// 	fmt.Println(i)
	// 	i++
	// }
	// i:=0
	// for{
	// 	fmt.Println(i)
	// 	i++
	// 	if i>=3{
	// 		break
	// 	}
	// }
	// i:=0
	//  for i<3{
	// 	if i==1{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	//
	// }
	nums := []int{1, 3, 2, 4, 0}
	for _, value := range nums {
		if value == 3 {
			break
		}
		fmt.Println(value)
	}
}
