package main

import (
	"fmt"
)

func main() {
	var emptyName string = ""
	var name string = "Pigi"

	fmt.Println(greetings(emptyName)) // this will return "Hello Dude!"
	fmt.Println(greetings(name)) // this will return "Hello Pigi!"
}