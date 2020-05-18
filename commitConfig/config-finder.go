package commitConfig

import (
	"log"
	"os/exec"
)

// GetGitRootDir blabla
func GetGitRootDir() string {
	myArguments := []string{"rev-parse", "--show-toplevel"}
	output, err := exec.Command("git", myArguments...).CombinedOutput()
	if err != nil {
		log.Fatal("Cannot find git root")
	}
	return string(output)
}
