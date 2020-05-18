package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ItsAnas/git-commit-configurator/commitConfig"
	"github.com/ItsAnas/git-commit-configurator/commiter"
	"github.com/ItsAnas/git-commit-configurator/interact"
)

func main() {

	path := commitConfig.GetGitRootDir()

	if path == "" {
		log.Fatal("Cannot find the root of git repository")
	}

	fmt.Println(path)
	os.Exit(0)

	if interact.AskForCommit() {
		commitType := interact.AskCommitType()
		commitMessage := interact.AskMessage()
		commiter.CommitMessage(commitType, commitMessage)
	}
}
