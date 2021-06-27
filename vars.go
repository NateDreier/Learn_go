package main

import "fmt"

func main() {
	var myInt int = 289
	var myFloat, ok, notOkay = 2.89, "yes", "no"

	fmt.Println("int:", myInt)
	fmt.Println("multi:", myInt * 25)
	fmt.Println("float:", myFloat)
	fmt.Println("yes:", ok)
	fmt.Println("notOk:", notOkay)
}