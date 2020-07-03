package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	greetings "github.com/Androide79/go-greetings"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// set config file path
	viper.SetConfigFile(".env")
	// read config file
	viper.ReadInConfig()

	// init logger function
	initLogger()

	var name string = viper.GetString("NAME")
	log.Infof("'name' variable init with '%v' value", name)
	var lang string = viper.GetString("LANGUAGE")
	log.Infof("'lang' variable init with '%v' value", lang)
	var covidURL string = viper.GetString("URL_COVID")

	fmt.Println(greetings.Greet(name, lang))

	// make an HTTP request
	response, err := http.Get(covidURL)
	if err != nil {
		// TODO: insert logging system to handle the prints
		fmt.Printf("An error occurred: %v", err)
	}

	defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	// TODO: insert logging system to handle the prints
	// 	fmt.Printf("An error occurred: %v", err)
	// }

	var covidData map[string]interface{}

	json.NewDecoder(response.Body).Decode(&covidData)

	fmt.Println(covidData["records"])
}

func initLogger() {
	// choose logger file
	file, err := os.OpenFile(viper.GetString("LOG_FILE_NAME"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// set log output to file
	log.SetOutput(file)
	// set the formatter (if you consider moving these data to a centralized system you can use 'JSONFormatter')
	log.SetFormatter(&log.TextFormatter{})
	// set log level
	log.SetLevel(log.InfoLevel)
}
