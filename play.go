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

	var name string = viper.GetString("NAME")
	var lang string = viper.GetString("LANGUAGE")

	fmt.Println(greetings.Greet(name, lang))
}
