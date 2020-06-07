package main

import "testing"

func TestHelloEmptyArgument(t *testing.T) {
	// test for empty argument
	var emptyName string = greetings("")

	if emptyName != "Hello Dude!" {
		t.Errorf("greetings(\"\") failed, expected %v, got %v", "Hello Dude!", emptyName)
	}
}

func TestHelloValidArgument(t *testing.T) {
	// test for a valid argument
	var name string = greetings("Pigi")

	if name != "Hello Pigi!" {
		t.Errorf("greetings(\"Pigi\") failed, expected %v, got %v", "Hello Pigi!", name)
	}
}
