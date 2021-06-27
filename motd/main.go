package main

import (
  "fmt"
  "motd/message"
)

func main() {
	m := message.Greeting("Nate", "hello")
	fmt.Println(m)
}

