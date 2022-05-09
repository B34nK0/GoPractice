package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Bean")
	fmt.Println(message)
}
