package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintln("Hi, %v. Welcome!", name)
	return message
}
