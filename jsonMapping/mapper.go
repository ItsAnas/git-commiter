package jsonMapping

import (
	"encoding/json"
	"fmt"
	"log"
)

type PrefixRule struct {
	Prefix      string
	Description string
}

type JsonConfig struct {
	Name        string
	Description string
	Rules       []PrefixRule
}

func EncodeJsonConfig(config JsonConfig) bool {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(string(jsonData))
	return true
}
