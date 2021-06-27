package main

import "fmt"

func main() {
        
	i := 0
	for i < 100 {
		if i % 3 == 0 && i % 5 == 0 {
			i++
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			i++
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			i++
			fmt.Println("Buzz")
		} else {
			i++
			fmt.Println(i)
		}
	}
}
