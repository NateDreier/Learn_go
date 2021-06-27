package message

import "fmt"

func Greeting(name, message string) (shalom string)  {
	shalom = fmt.Sprintf("%s, %s, %s, %s", message, name, name, message)
	return
}