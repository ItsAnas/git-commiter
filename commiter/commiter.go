package commiter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// ParseCommitMessage add the message Type following the real user message
func parseCommitMessage(commitType string, input string) string {
	return fmt.Sprintf("%s: %s", commitType, input)
}

// CommitMessage launch command to commit the message
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
