package commiter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// ParseCommitMessage fekopkp
func parseCommitMessage(input string) string {
	return fmt.Sprintf("\"feat(git): %s\"", input)
}

// CommitMessage should
func CommitMessage(input string) bool {
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
