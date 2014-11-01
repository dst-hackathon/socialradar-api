package configuration

import (
	"encoding/json"
	"os"
	"fmt"
	"log"
)

type Configuration struct {
	DbHost		string
	DbPort		string
	DbUser		string
	DbPassword	string
	AvatarPath 	string
	ApiSignKey  string
}

func ReadFile() Configuration {
	file, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/dst-hackathon/socialradar-api/config.json")

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)

	if err != nil {
	  fmt.Println("error:", err)
	}

	fmt.Println(configuration)
	return configuration
}
