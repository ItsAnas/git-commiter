package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ItsAnas/git-commit-configurator/commitConfig"
	"github.com/ItsAnas/git-commit-configurator/commiter"
	"github.com/ItsAnas/git-commit-configurator/interact"
	"github.com/ItsAnas/git-commit-configurator/jsonMapping"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	config := jsonMapping.DecodeJsonConfig(".commit.json")
	jsonMapping.EncodeJsonConfig(config)

	gitRootDir := commitConfig.GetGitRootDir()
	configPath, err := commitConfig.FindCommitConfig(gitRootDir)

	if err != nil || configPath == "" {
		log.Fatal("Cannot find .commit.json")
	}

	fmt.Println(configPath)
	os.Exit(0)

	if interact.AskForCommit() {
		commitType := interact.AskCommitType()
		commitMessage := interact.AskMessage()
		commiter.CommitMessage(commitType, commitMessage)
	}

}
