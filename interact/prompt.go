package interact

import (
	"errors"
	"log"
	"unicode/utf8"

	"github.com/manifoldco/promptui"
)

// AskForCommit wesh
func AskForCommit() bool {
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

// AskForCommit wesh
func AskCommitType() string {
	prompt := promptui.Select{
		Label: "Type of commit",
		Items: []string{"feat", "fix", "ci", "tests", "doc"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}

// AskMessage is blabla
func AskMessage() string {

	validate := func(input string) error {
		if utf8.RuneCountInString(input) < 4 {
			return errors.New("Please more than 4 character")
		}

		if utf8.RuneCountInString(input) > 50 {
			return errors.New("Commit cannot exceed 50 characters")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter your commit message",
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}
