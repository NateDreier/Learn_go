package main

import "fmt"

func main() {
	metasyntatical := map[string]string{
		"foo": "bar",
		"baz": "quuz",
	}
	foo := map[string]int{}

	foo["bar"] = 1
	foo["baz"] = 2

	delete(foo, "bar")

	fmt.Println("list:", metasyntatical)
	fmt.Println("foo:", foo)
	fmt.Println(metasyntatical["foo"])
}