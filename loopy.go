package main

import "fmt"

func main() {
	fmt.Println("shalom")

	foo := map[string]int{}
	foo["bar"] = 1
	foo["baz"] = 2
	foo["quux"] = 3

	for metasyn, val := range foo {
		fmt.Println(fmt.Sprintf("something something %s something something %v", metasyn, val))
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	/*
	infinite loop
	a := 0
	for a < 10 {
		fmt.Println(a)
	}
	*/

	a := 0
	for a < 10 {
		if a % 2 == 0 {
			a ++
			continue
		} else if a == 5 {
			break
		} 
		
		fmt.Println("we counting")
		a++
	}
}