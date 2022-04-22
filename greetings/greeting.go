package greetings

import "fmt"

func hello(name string) string {
	message := fmt.Sprint("Hi, Welcome", name)
	return message
}