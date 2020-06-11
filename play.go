package main

import (
	"fmt"
	greetings "github.com/Androide79/go-greetings"
)

func main() {
	var name string = "Pigi"

	fmt.Println(greetings.Greet(name))
}