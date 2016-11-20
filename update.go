package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func update(c *cli.Context) error {
	spinner := NewSpinner("Updating...").Start()

	repos := Config

	updatedPaths := []string{}

	for _, repo := range repos {
		path := deducePath(repo)
		if matched, err := regexp.MatchString("^/", path); err == nil && !matched {
			path = rootDir + "/" + path
		}

		updatedPaths = append(updatedPaths, updateRepo(path, repo))
	}

	spinner.Stop()

	for _, updatedPath := range updatedPaths {
		fmt.Printf("Updated the following: %s", updatedPath)
	}

	return nil

}

func updateRepo(path string, repo Dot) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if v, ok := repo["repo"]; ok {
			cmdArgs := []string{
				"clone", "--depth", "1", "--recursive", v, path,
			}

			fmt.Printf("Args %v\n", cmdArgs)

			out, err := exec.Command("git", cmdArgs...).Output()
			if err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf("Output %v\n", out)
			return path
		} else {
			log.Fatalf("There is no repo entry!")
		}
	} else {
		if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
			log.Fatalf("The path exists but is not a git repo")
		} else {
			cmd := fmt.Sprintf("cd %s && git pull --ff-only && git submodule update --init;", path)
			_, err := exec.Command("/bin/sh", "-c", cmd).Output()
			if err != nil {
				log.Fatal(err)
			}
			return path
		}
	}
	return ""

}

func deducePath(repo Dot) string {
	if v, ok := repo["path"]; ok {
		return v
	}
	return ""
}
