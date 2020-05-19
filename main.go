package main

import (
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
	gitRootDir := commitConfig.GetGitRootDir()
	configPath, err := commitConfig.FindCommitConfig(gitRootDir)

	if err != nil || configPath == "" {
		if !interact.AskForCreate() {
			os.Exit(3)
		}
		jsonMapping.CreateSample()
	}

	if interact.AskForCommit() {
		config := jsonMapping.DecodeJsonConfig(configPath)
		commitType := interact.AskCommitType(config)
		commitMessage := interact.AskMessage()
		commiter.CommitMessage(commitType, commitMessage)
	}
}
