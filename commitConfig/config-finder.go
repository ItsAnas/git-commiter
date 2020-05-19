package commitConfig

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

// GetGitRootDir returns the root of the git repository
func GetGitRootDir() string {
	myArguments := []string{"rev-parse", "--show-toplevel"}
	output, err := exec.Command("git", myArguments...).CombinedOutput()
	if err != nil {
		log.Fatal("Cannot find git repository")
	}
	return string(output)
}

func removeEol(input string) string {
	lastCharacter := input[len(input)-1:]
	fmt.Printf("last char is %s\n", lastCharacter)
	if lastCharacter == "\n" {
		return input[len(input)-1:]
	}
	return input
}

// FindCommitConfig returns the path of the .commit.json file
// if it exists or an error if not
func FindCommitConfig(rootPath string) (string, error) {
	rootPath = strings.TrimSuffix(rootPath, "\n")
	filePath := path.Join(rootPath, ".commit.json")
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return "", errors.New("Cannot find .commit.json")
	}

	return filePath, nil
}
