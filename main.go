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

func main() {

	config := jsonMapping.JsonConfig{
		Name:        "My Config",
		Description: "Hey this is my config for commit",
		Rules: []jsonMapping.PrefixRule{
			jsonMapping.PrefixRule{
				Prefix:      "feat",
				Description: "feat: Implement new feature",
			},
			jsonMapping.PrefixRule{
				Prefix:      "doc",
				Description: "doc: writing doc",
			},
			jsonMapping.PrefixRule{
				Prefix:      "fix",
				Description: "fix: fix bug",
			},
		},
	}

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
