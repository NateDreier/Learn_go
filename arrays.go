package main

import "fmt"

func main() {
	// use make if you know the length
	slice_make := make([]string, 4)
	slice_make[0] = "foo"
	slice_make[1] = "bar"
	slice_make[2] = "bazinga"
	slice_make[3] = "sup"

	// slice incoming
	slicer := []string{}
	slicer = append(slicer, "foo")
	name := [3]string{"foo", "bar", "baz"}

	var names [3]string

	names[0] = "bazing"
	names[1] = "fazinga"
	//names[2] = "barzinga"

	fmt.Println(name)
	fmt.Println(names)
	fmt.Println("empty name", names[2] == "")
	fmt.Println("pring slicer:", slicer)
	fmt.Println("slicer_make", slice_make)
}