package main

import "fmt"

func greetings(name string) string {
	var result string

	if len(name) == 0 {
		result = "Hello Dude!"
	} else {
		result = fmt.Sprintf("Hello %v!", name)
	}

	return result
}
