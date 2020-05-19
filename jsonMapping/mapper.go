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

func EncodeJsonConfig(config CommitConfig) []byte {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(string(jsonData))
	return jsonData
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

func CreateSample() bool {
	config := CommitConfig{
		Name:        "My Config",
		Description: "Hey this is my config for commit",
		Rules: []PrefixRule{
			PrefixRule{
				Prefix:      "feat",
				Description: "feat: Implement new feature"},
			PrefixRule{
				Prefix:      "doc",
				Description: "doc: writing doc",
			},
			PrefixRule{
				Prefix:      "fix",
				Description: "fix: fix bug",
			},
		},
	}

	bytes := EncodeJsonConfig(config)

	if bytes == nil {
		log.Fatal("Cannot decode config. Please provide an issue on github")
	}

	f, err := os.Create(".commit.json")
	if err != nil {
		log.Fatal("Cannot create .commit.json")
	}

	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		log.Fatal("Cannot create .commit.json")
	}

	return true
}
