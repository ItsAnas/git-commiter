package interact

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/ItsAnas/git-commit-configurator/jsonMapping"
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

func getDescriptions(config jsonMapping.CommitConfig) []string {
	var res []string
	for _, rule := range config.Rules {
		res = append(res, rule.Description)
	}
	return res
}

// AskForCommit wesh
func AskCommitType(config jsonMapping.CommitConfig) string {
	descriptions := getDescriptions(config)
	prompt := promptui.Select{
		Label: "Type of commit",
		Items: descriptions,
	}
	index, _, err := prompt.Run()
	fmt.Printf("value choosen is: %d", index)
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return config.Rules[index].Prefix
}

// AskForCommit wesh
func AskForCreate() bool {
	prompt := promptui.Select{
		Label: "No .commit.json found. Do you want to create one ? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
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
