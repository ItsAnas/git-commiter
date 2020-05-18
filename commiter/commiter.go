package commiter

import "fmt"

// ParseCommitMessage fekopkp
func ParseCommitMessage(input string) string {
	return fmt.Sprintf("\"feat(git): %s\"", input)
}
