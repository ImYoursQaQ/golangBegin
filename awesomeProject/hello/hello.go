package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	// 0 : 无， 1 ： date 2 ： time 3 ： date + time
	log.SetFlags(3)
	//message := greetings.Hello("teo")
	//fmt.Println(message)

	//message, err := greetings.HelloWithError("teo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(message)

	//message, err := greetings.HelloWithRandom("teo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(message)

	names := []string{
		"zhang",
		"li",
		"wang",
	}

	messages, err := greetings.Hellos(names)
	if err == nil {
		fmt.Println(messages)
	}
}
