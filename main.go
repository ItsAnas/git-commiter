package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello, world.")
	myArguments := []string{"branch", "--all"}
	output, err := exec.Command("git", myArguments...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}
