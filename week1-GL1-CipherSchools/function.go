package main

import "fmt"

/*
func function_name(){
	statement-1
	statement-2
	statement-3
	statement-4
}
*/
func main() {
	//result, x := a()
	w := new(int)
	name := new(string)
	result, x := c(w, name)
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(*name)
	r, _ := b(1, 2, 3, 4, 5, 6, 6)
	fmt.Println(r)
}
func a() (int, string) {
	return 122, "weqrr3"
}
func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c(w *int, name *string) (i int, j string) {
	i = 10
	j = "rahul"
	*w = 100
	*name = "code"
	return
}
