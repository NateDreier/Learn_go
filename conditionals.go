package main

import "fmt"

func main() {
	foo := map[string]int{}
	foo["bar"] = 102
	foo["baz"] = 20
	foo["quuz"] = 30

	switch foo["baz"] {
	case 20:
		fmt.Println("baz == 20")
	case 21:
		fmt.Println("baz can drink")
	default:
		fmt.Println("hmm")		
	}

	switch {
	case foo["bar"] < 20:
		fmt.Println("bar is less than 20")
	case foo["bar"] > 20:
		fmt.Println("bar is greater than 20")
	default:
		fmt.Println("something aint right")
	}

	if foo["bar"] < 20 {
		fmt.Println("bar is less than 20")
	} else if foo["bar"] > 20 {
		fmt.Println("ya boi is greater than  20")
	} else {
		fmt.Println("something aint right")
	}
}