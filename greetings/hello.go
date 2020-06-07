package main

import "fmt"

func greetings (name string) string {
	if( len(name) == 0) {
		return "Hello Dude!"
	} else {
		return fmt.Sprintf("Hello %v!", name)
	}
}