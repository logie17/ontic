package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/exec"
)

var (
	scriptPath        = rootDir + "/lib"
	confFile          = rootDir + "/conf.json"
	onticShellGitPath = "https://github.com/logie17/ontic-shell"
)

func initOntic(c *cli.Context) error {
	// what do we want to do with init
	// 1. check if .ontic exits, if not create it
	// 2. Check if ./ontic/conf.json exists, if not, create it
	spinner := NewSpinner("Updating...").Start()

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		os.Mkdir(rootDir, 0700)
	}

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		f, errF := os.Create(confFile)
		if errF != nil {
			log.Fatalf("Do something with error %v", errF)
		}
		defer f.Close()

		f.WriteString(defaultConf())
		f.Sync()

	}

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

func defaultConf() string {
	return `
{
    "dots": [
	{
	    "path": "loop-dots",
	    "repo": "git@github.com:logie17/loop-dots.git"
	}
    ]
}
`
}
