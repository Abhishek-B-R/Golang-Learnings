package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Abhishek")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(message)


    names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(messages)
}