package commiter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// ParseCommitMessage fekopkp
func parseCommitMessage(commitType string, input string) string {
	return fmt.Sprintf("%s: %s", commitType, input)
}

// CommitMessage should
func CommitMessage(commitType string, input string) bool {
	formattedMessage := parseCommitMessage(commitType, input)
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
