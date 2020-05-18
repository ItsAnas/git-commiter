package main

import (
	"github.com/ItsAnas/git-commit-configurator/commiter"
	"github.com/ItsAnas/git-commit-configurator/interact"
)

func main() {
	if interact.AskForCommit() {
		commitType := interact.AskCommitType()
		commitMessage := interact.AskMessage()
		commiter.CommitMessage(commitType, commitMessage)
	}
}
