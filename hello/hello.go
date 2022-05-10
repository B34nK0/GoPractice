package main

import (
	"fmt"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("gretings ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		//打印错误以及停止程序
		log.Fatal(err)
	}
	fmt.Println(messages)

	message, err := greetings.Hello("")
	if err != nil {
		//打印错误以及停止程序
		log.Fatal(err)
	}
	fmt.Println(message)
}
