package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/exec"
)

var (
	scriptPath        = rootDir + "/scripts"
	onticShellGitPath = "https://github.com/logie17/ontic-shell"
)

func initOntic(c *cli.Context) error {
	spinner := NewSpinner("Updating...").Start()

	os.Mkdir(rootDir, 0700)
	cmdArgs := []string{
		"clone", "--depth", "1", "--recursive", onticShellGitPath, scriptPath,
	}

	out, err := exec.Command("git", cmdArgs...).Output()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Output %v\n", out)

	spinner.Stop()
	return nil
}
