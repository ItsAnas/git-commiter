package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func askForCommit() bool {
	prompt := promptui.Select{
		Label: "Do you want to commit ? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func parseCommitMessage(input string) string {
	return fmt.Sprintf("\"feat(git): %s\"", input)
}

func commitMessage(input string) bool {
	formattedMessage := parseCommitMessage(input)
	myArguments := []string{"commit", "-m", formattedMessage}
	output, err := exec.Command("git", myArguments...).CombinedOutput()
	if err != nil {
		log.Fatal("Error during commit")
		os.Stderr.WriteString(err.Error())
		return false
	}
	fmt.Println(string(output))
	return true
}

func main() {
	if askForCommit() {
		commitMessage("add error message when commit fail")
	}
}
