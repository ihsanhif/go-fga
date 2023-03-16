package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true
	var k float64 = 123.456
	ya := 'Я'

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	fmt.Printf("%t \n", j)
	fmt.Printf("%#X \n", ya)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", i)
	fmt.Printf("%U \n", ya)
	fmt.Printf("%.3f \n", k)
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}
