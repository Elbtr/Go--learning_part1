package main

import "fmt"

// operasi matematika

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 20
	var b = 2

	var c = a * b 
	fmt.Println(c)

	// contoh Augmented Assignments

	var i = 10
	i += 10
	fmt.Println(i)


	// contoh Unary Operator

	i++ 
	fmt.Println(i)

	var negative = -100
	var positive = +100

	fmt.Println(negative)
	fmt.Println(positive)


}

