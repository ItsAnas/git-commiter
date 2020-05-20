package main

import (
	"os"

	"github.com/ItsAnas/git-commiter/commitConfig"
	"github.com/ItsAnas/git-commiter/commiter"
	"github.com/ItsAnas/git-commiter/interact"
	"github.com/ItsAnas/git-commiter/jsonMapping"
)

func main() {
	gitRootDir := commitConfig.GetGitRootDir()
	configPath, err := commitConfig.FindCommitConfig(gitRootDir)

	if err != nil || configPath == "" {
		if !interact.AskForCreate() {
			os.Exit(3)
		}
		jsonMapping.CreateSample()
		configPath, _ = commitConfig.FindCommitConfig(gitRootDir)
	}

	if interact.AskForCommit() {
		config := jsonMapping.DecodeJsonConfig(configPath)
		commitType := interact.AskCommitType(config)
		commitMessage := interact.AskMessage()
		commiter.CommitMessage(commitType, commitMessage)
	}
}
