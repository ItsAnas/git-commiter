package jsonMapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type PrefixRule struct {
	Prefix      string
	Description string
}

type CommitConfig struct {
	Name        string
	Description string
	Rules       []PrefixRule
}

func EncodeJsonConfig(config CommitConfig) bool {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(string(jsonData))
	return true
}

func DecodeJsonConfig(input string) CommitConfig {

	// Open our jsonFile
	jsonFile, err := os.Open(input)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", input)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	jsonData, _ := ioutil.ReadAll(jsonFile)
	var config CommitConfig
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Println(err)
	}
	return config
}
