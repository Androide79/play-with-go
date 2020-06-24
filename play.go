package main

import (
	"fmt"

	greetings "github.com/Androide79/go-greetings"
	"github.com/spf13/viper"
)

func main() {
	// set config file path
	viper.SetConfigFile(".env")
	// read config file
	viper.ReadInConfig()

	fmt.Println(">>> ", viper.GetString("NAME"))

	var name string = "Pigi"

	fmt.Println(greetings.Greet(name))
}
