package main

import (
  "fmt"
  "motd/message"
  "bufio"
//  "io/ioutil"
  "os"
  "strings"
)

func main() {
	f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
	
	if err != nil {
		fmt.Println("err err err")
		os.Exit(1)
	}

	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	_, err = f.Write([]byte(m))

	//err := ioutil.WriteFile("/etc/motd", []byte(m), 0644)

	if err != nil {
		fmt.Println("can't write!!")
		os.Exit(1)
	}

}

