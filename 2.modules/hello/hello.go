package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	message, err := greetings.Hello("Abhishek")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
}