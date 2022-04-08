package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Pablo")
	fmt.Println(message)
}
