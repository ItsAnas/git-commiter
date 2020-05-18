package main

import (
	"fmt"

	"github.com/ItsAnas/git-commit-configurator/commiter"
	"github.com/ItsAnas/git-commit-configurator/interact"
)

func main() {
	fmt.Println("It works ! 2")
	if interact.AskForCommit() {
		commiter.CommitMessage("add error message when commit fail")
	}
}
